package repository

import (
	"slices"

	"cyskillswap/internal/constants"
	bizerr "cyskillswap/internal/errors"
	"cyskillswap/internal/model"
)

func GetOrderByID(id int) (model.ExchangeOrder, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()
	for _, order := range db.orders {
		if order.ID == id {
			order.Confirmations = slices.Clone(order.Confirmations)
			return order, true
		}
	}
	return model.ExchangeOrder{}, false
}

func ListOrders() []model.ExchangeOrder {
	db.mu.Lock()
	defer db.mu.Unlock()
	out := make([]model.ExchangeOrder, len(db.orders))
	for i, order := range db.orders {
		order.Confirmations = slices.Clone(order.Confirmations)
		out[i] = order
	}
	return out
}

// HasConfirmedOrderAt 判断用户在某时段是否已有确认预约（冲突检测的唯一事实来源）
func HasConfirmedOrderAt(user string, timeSlot string, excludeOrderID int) bool {
	db.mu.Lock()
	defer db.mu.Unlock()
	return hasConfirmedOrderAtLocked(user, timeSlot, excludeOrderID)
}

func hasConfirmedOrderAtLocked(user string, timeSlot string, excludeOrderID int) bool {
	for _, order := range db.orders {
		if order.ID == excludeOrderID || order.Status != constants.OrderStatusConfirmed {
			continue
		}
		if order.TimeSlot == timeSlot && order.Involves(user) {
			return true
		}
	}
	return false
}

// ApplySelection 原子完成"选择响应"：防御性复检后生成交换单，
// 选中响应置为已选中，其余响应置为未选中，需求进入匹配中
func ApplySelection(needID int, responseID int, timeSlot string) (model.ExchangeOrder, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	needIdx := -1
	for i := range db.needs {
		if db.needs[i].ID == needID {
			needIdx = i
		}
	}
	if needIdx < 0 {
		return model.ExchangeOrder{}, bizerr.ErrNeedNotFound
	}
	need := db.needs[needIdx]
	if need.Status != constants.NeedStatusOpen {
		return model.ExchangeOrder{}, bizerr.ErrNeedClosed
	}

	respIdx := -1
	for i := range db.responses {
		if db.responses[i].ID == responseID && db.responses[i].NeedID == needID {
			respIdx = i
		}
	}
	if respIdx < 0 {
		return model.ExchangeOrder{}, bizerr.ErrResponseNotFound
	}
	resp := db.responses[respIdx]
	if resp.Status != constants.ResponseStatusPending {
		return model.ExchangeOrder{}, bizerr.ErrResponseNotPending
	}

	// 同一需求只允许存在一张交换单，防止重复生成
	for _, order := range db.orders {
		if order.NeedID == needID {
			return model.ExchangeOrder{}, bizerr.ErrOrderExists
		}
	}

	order := model.ExchangeOrder{
		ID:            db.nextOrderID,
		NeedID:        needID,
		NeedTitle:     need.Title,
		ResponseID:    responseID,
		Requester:     need.Requester,
		Responder:     resp.Responder,
		OfferSkill:    resp.OfferSkill,
		TimeSlot:      timeSlot,
		Mode:          resp.Mode,
		Place:         resp.Place,
		Agenda:        need.Title + " · " + resp.OfferSkill,
		Status:        constants.OrderStatusPending,
		Confirmations: []string{},
		CreatedAt:     NowString(),
	}
	db.nextOrderID++
	db.orders = append(db.orders, order)

	for i := range db.responses {
		if db.responses[i].NeedID != needID {
			continue
		}
		if db.responses[i].ID == responseID {
			db.responses[i].Status = constants.ResponseStatusSelected
		} else {
			db.responses[i].Status = constants.ResponseStatusNotSelected
		}
	}
	db.needs[needIdx].Status = constants.NeedStatusMatching

	order.Confirmations = slices.Clone(order.Confirmations)
	return order, nil
}

// ApplyConfirmation 原子完成参与者确认：复检身份与冲突后记录确认，
// 双方均确认后交换单生效、需求标记为已约成
func ApplyConfirmation(orderID int, user string) (model.ExchangeOrder, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	idx := -1
	for i := range db.orders {
		if db.orders[i].ID == orderID {
			idx = i
		}
	}
	if idx < 0 {
		return model.ExchangeOrder{}, bizerr.ErrOrderNotFound
	}
	order := &db.orders[idx]
	if order.Status != constants.OrderStatusPending {
		return model.ExchangeOrder{}, bizerr.ErrOrderNotPending
	}
	if !order.Involves(user) {
		return model.ExchangeOrder{}, bizerr.ErrNotParticipant
	}
	if order.ConfirmedBy(user) {
		return model.ExchangeOrder{}, bizerr.ErrAlreadyConfirmed
	}
	if hasConfirmedOrderAtLocked(user, order.TimeSlot, order.ID) {
		return model.ExchangeOrder{}, bizerr.ScheduleConflict(order.TimeSlot)
	}

	order.Confirmations = append(order.Confirmations, user)
	if len(order.Confirmations) == 2 {
		order.Status = constants.OrderStatusConfirmed
		for i := range db.needs {
			if db.needs[i].ID == order.NeedID {
				db.needs[i].Status = constants.NeedStatusDone
			}
		}
	}

	confirmed := *order
	confirmed.Confirmations = slices.Clone(order.Confirmations)
	return confirmed, nil
}
