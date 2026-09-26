package repository

import (
	"slices"

	"cyskillswap/internal/constants"
	"cyskillswap/internal/model"
)

func ListSkills() []model.Skill { return db.skills }

func ListNeeds() []model.Need {
	db.mu.Lock()
	defer db.mu.Unlock()
	return slices.Clone(db.needs)
}

func ListMatches() []model.Match { return db.matches }

// ListAppointments 预约时间线由已确认的交换单派生，保证与冲突检测同源
func ListAppointments() []model.Appointment {
	db.mu.Lock()
	defer db.mu.Unlock()
	appointments := make([]model.Appointment, 0, len(db.orders))
	for _, order := range db.orders {
		if order.Status != constants.OrderStatusConfirmed {
			continue
		}
		appointments = append(appointments, model.Appointment{
			ID:     order.ID,
			Pair:   order.Requester + " ↔ " + order.Responder,
			Time:   order.TimeSlot,
			Mode:   order.Mode,
			Place:  order.Place,
			Status: constants.AppointmentStatusConfirmed,
			Agenda: order.Agenda,
		})
	}
	return appointments
}

func ListReviews() []model.Review { return db.reviews }

func ListMessages() []model.Conversation { return db.messages }

func GetProfile() model.Profile {
	return model.Profile{
		Name: "林澈", Major: "新闻传播 2023", CreditScore: 91, CreditLevel: constants.CreditGold,
		SkillWall: ListSkills()[:1],
		Radar:     map[string]int{"摄影": 92, "修图": 86, "沟通": 90, "编程": 42, "乐器": 35},
		History:   []string{"完成毕业照拍摄交换", "响应 Python 数据分析需求", "预约吉他入门课"},
		Reviews:   ListReviews(),
	}
}
