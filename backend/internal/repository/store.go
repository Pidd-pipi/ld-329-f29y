package repository

import (
	"sync"

	"cyskillswap/internal/constants"
	"cyskillswap/internal/errors"
	"cyskillswap/internal/model"
)

// Store 是内存数据仓库，所有读写通过互斥锁保证并发安全。
type Store struct {
	mu           sync.RWMutex
	skills       []model.Skill
	needs        []model.Need
	responses    []model.Response
	orders       []model.SwapOrder
	matches      []model.Match
	appointments []model.Appointment
	reviews      []model.Review
	messages     []model.Conversation

	nextResponseID    int
	nextOrderID       int
	nextAppointmentID int
}

func NewStore() *Store {
	responses := seedResponses()
	s := &Store{
		skills:       seedSkills(),
		needs:        seedNeeds(),
		responses:    responses,
		orders:       []model.SwapOrder{},
		matches:      seedMatches(),
		appointments: seedAppointments(),
		reviews:      seedReviews(),
		messages:     seedMessages(),

		nextResponseID:    len(responses) + 1,
		nextOrderID:       1,
		nextAppointmentID: len(seedAppointments()) + 1,
	}
	return s
}

var defaultStore = NewStore()

func ListSkills() []model.Skill             { return defaultStore.ListSkills() }
func ListMatches() []model.Match            { return defaultStore.ListMatches() }
func ListReviews() []model.Review           { return defaultStore.ListReviews() }
func ListMessages() []model.Conversation    { return defaultStore.ListMessages() }
func ListAppointments() []model.Appointment { return defaultStore.ListAppointments() }
func ListNeeds() []model.Need               { return defaultStore.ListNeeds() }
func GetNeed(id int) (model.Need, bool)     { return defaultStore.GetNeed(id) }
func ListResponses(needID int) []model.Response {
	return defaultStore.ListResponses(needID)
}
func ListSwapOrders() []model.SwapOrder { return defaultStore.ListSwapOrders() }
func GetProfile() model.Profile         { return defaultStore.GetProfile() }
func CreateResponse(needID int, input model.ResponseInput) (model.Response, error) {
	return defaultStore.CreateResponse(needID, input)
}
func AcceptResponse(needID, responseID int) (model.SwapOrder, error) {
	return defaultStore.AcceptResponse(needID, responseID)
}
func ConfirmSwapOrder(id int, user string) (model.SwapOrder, error) {
	return defaultStore.ConfirmSwapOrder(id, user)
}

func (s *Store) ListSkills() []model.Skill {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]model.Skill(nil), s.skills...)
}

func (s *Store) ListMatches() []model.Match {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]model.Match(nil), s.matches...)
}

func (s *Store) ListReviews() []model.Review {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]model.Review(nil), s.reviews...)
}

func (s *Store) ListMessages() []model.Conversation {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]model.Conversation(nil), s.messages...)
}

func (s *Store) ListAppointments() []model.Appointment {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]model.Appointment(nil), s.appointments...)
}

// ListNeeds 返回需求列表，响应数量按实际响应记录实时统计。
func (s *Store) ListNeeds() []model.Need {
	s.mu.RLock()
	defer s.mu.RUnlock()
	needs := append([]model.Need(nil), s.needs...)
	for i := range needs {
		count := 0
		for _, resp := range s.responses {
			if resp.NeedID == needs[i].ID {
				count++
			}
		}
		needs[i].Responses = count
	}
	return needs
}

func (s *Store) GetNeed(id int) (model.Need, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, need := range s.needs {
		if need.ID == id {
			return need, true
		}
	}
	return model.Need{}, false
}

func (s *Store) ListResponses(needID int) []model.Response {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := []model.Response{}
	for _, resp := range s.responses {
		if resp.NeedID == needID {
			result = append(result, resp)
		}
	}
	return result
}

func (s *Store) ListSwapOrders() []model.SwapOrder {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]model.SwapOrder(nil), s.orders...)
}

// GetProfile 返回个人主页，并附带与本人相关的交换单和响应记录。
func (s *Store) GetProfile() model.Profile {
	s.mu.RLock()
	defer s.mu.RUnlock()
	profile := baseProfile()
	profile.SkillWall = append([]model.Skill(nil), s.skills[:1]...)
	profile.Reviews = append([]model.Review(nil), s.reviews...)
	profile.Orders = []model.SwapOrder{}
	for _, order := range s.orders {
		if order.Requester == profile.Name || order.Responder == profile.Name {
			profile.Orders = append(profile.Orders, order)
		}
	}
	profile.MyResponses = []model.Response{}
	for _, resp := range s.responses {
		if resp.Responder == profile.Name {
			profile.MyResponses = append(profile.MyResponses, resp)
		}
	}
	return profile
}

// CreateResponse 校验需求状态后新增一条等待中的响应。
func (s *Store) CreateResponse(needID int, input model.ResponseInput) (model.Response, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	needIdx := s.findNeedIndex(needID)
	if needIdx < 0 {
		return model.Response{}, errors.NeedNotFound()
	}
	need := s.needs[needIdx]
	if need.Status != constants.NeedStatusOpen {
		return model.Response{}, errors.NeedClosed()
	}
	if need.Requester == input.Responder {
		return model.Response{}, errors.SelfResponse()
	}
	for _, resp := range s.responses {
		if resp.NeedID == needID && resp.Responder == input.Responder {
			return model.Response{}, errors.DuplicateResponse()
		}
	}
	resp := model.Response{
		ID:         s.nextResponseID,
		NeedID:     needID,
		Responder:  input.Responder,
		OfferSkill: input.OfferSkill,
		TimeSlot:   input.TimeSlot,
		PlaceType:  input.PlaceType,
		Place:      input.Place,
		Note:       input.Note,
		Status:     constants.ResponseStatusPending,
	}
	s.nextResponseID++
	s.responses = append(s.responses, resp)
	return resp, nil
}

// AcceptResponse 由发布者选中一条等待中的响应：
// 先校验时段冲突与重复交换单，通过后生成待双方确认的交换单，
// 选中的响应标为已选中，其余响应标为未选中，需求进入已约成。
func (s *Store) AcceptResponse(needID, responseID int) (model.SwapOrder, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	needIdx := s.findNeedIndex(needID)
	if needIdx < 0 {
		return model.SwapOrder{}, errors.NeedNotFound()
	}
	if s.needs[needIdx].Status != constants.NeedStatusOpen {
		return model.SwapOrder{}, errors.NeedClosed()
	}
	respIdx := s.findResponseIndex(needID, responseID)
	if respIdx < 0 {
		return model.SwapOrder{}, errors.ResponseNotFound()
	}
	resp := s.responses[respIdx]
	if resp.Status != constants.ResponseStatusPending {
		return model.SwapOrder{}, errors.ResponseNotPending()
	}
	for _, order := range s.orders {
		if order.NeedID == needID {
			return model.SwapOrder{}, errors.DuplicateOrder()
		}
	}
	requester := s.needs[needIdx].Requester
	for _, participant := range []string{requester, resp.Responder} {
		if s.hasConfirmedAppointmentLocked(participant, resp.TimeSlot) {
			return model.SwapOrder{}, errors.ScheduleConflict(participant, resp.TimeSlot)
		}
	}
	order := model.SwapOrder{
		ID:         s.nextOrderID,
		NeedID:     needID,
		NeedTitle:  s.needs[needIdx].Title,
		ResponseID: resp.ID,
		Requester:  requester,
		Responder:  resp.Responder,
		OfferSkill: resp.OfferSkill,
		TimeSlot:   resp.TimeSlot,
		Place:      resp.Place,
		Status:     constants.SwapOrderStatusPending,
	}
	s.nextOrderID++
	s.orders = append(s.orders, order)
	for i := range s.responses {
		if s.responses[i].NeedID != needID {
			continue
		}
		if s.responses[i].ID == resp.ID {
			s.responses[i].Status = constants.ResponseStatusChosen
		} else {
			s.responses[i].Status = constants.ResponseStatusRejected
		}
	}
	s.needs[needIdx].Status = constants.NeedStatusBooked
	return order, nil
}

// ConfirmSwapOrder 由交换双方之一确认交换单，双方都确认后生效并生成预约。
// 生效前再次校验同时段确认预约，冲突则保持待确认并提示。
func (s *Store) ConfirmSwapOrder(id int, user string) (model.SwapOrder, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	orderIdx := -1
	for i := range s.orders {
		if s.orders[i].ID == id {
			orderIdx = i
			break
		}
	}
	if orderIdx < 0 {
		return model.SwapOrder{}, errors.OrderNotFound()
	}
	order := &s.orders[orderIdx]
	if order.Status == constants.SwapOrderStatusConfirmed {
		return model.SwapOrder{}, errors.AlreadyConfirmed()
	}
	var side *bool
	switch user {
	case order.Requester:
		side = &order.RequesterConfirmed
	case order.Responder:
		side = &order.ResponderConfirmed
	default:
		return model.SwapOrder{}, errors.NotParticipant()
	}
	if *side {
		return model.SwapOrder{}, errors.AlreadyConfirmed()
	}
	completing := (user == order.Requester && order.ResponderConfirmed) ||
		(user == order.Responder && order.RequesterConfirmed)
	if completing {
		for _, participant := range []string{order.Requester, order.Responder} {
			if s.hasConfirmedAppointmentLocked(participant, order.TimeSlot) {
				return model.SwapOrder{}, errors.ScheduleConflict(participant, order.TimeSlot)
			}
		}
	}
	*side = true
	if order.RequesterConfirmed && order.ResponderConfirmed {
		order.Status = constants.SwapOrderStatusConfirmed
		s.appointments = append(s.appointments, model.Appointment{
			ID:           s.nextAppointmentID,
			Pair:         order.Requester + " ↔ " + order.Responder,
			Time:         order.TimeSlot,
			TimeSlot:     order.TimeSlot,
			Place:        order.Place,
			Status:       constants.AppointmentStatusConfirmed,
			Agenda:       order.NeedTitle + " · " + order.OfferSkill,
			Participants: []string{order.Requester, order.Responder},
		})
		s.nextAppointmentID++
	}
	return *order, nil
}

func (s *Store) findNeedIndex(id int) int {
	for i := range s.needs {
		if s.needs[i].ID == id {
			return i
		}
	}
	return -1
}

func (s *Store) findResponseIndex(needID, responseID int) int {
	for i := range s.responses {
		if s.responses[i].ID == responseID && s.responses[i].NeedID == needID {
			return i
		}
	}
	return -1
}

// hasConfirmedAppointmentLocked 判断用户在某时段是否已有双方已确认的预约，调用前必须持有锁。
func (s *Store) hasConfirmedAppointmentLocked(user, slot string) bool {
	for _, appt := range s.appointments {
		if appt.Status != constants.AppointmentStatusConfirmed || appt.TimeSlot != slot {
			continue
		}
		for _, participant := range appt.Participants {
			if participant == user {
				return true
			}
		}
	}
	return false
}
