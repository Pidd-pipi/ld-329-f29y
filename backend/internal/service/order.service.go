package service

import (
	"slices"
	"strings"

	"cyskillswap/internal/constants"
	bizerr "cyskillswap/internal/errors"
	"cyskillswap/internal/model"
	"cyskillswap/internal/repository"
)

// SelectResponse 发布者从等待中的响应里选一位，生成双方待确认的交换单
func SelectResponse(needID int, in model.SelectResponseInput) (model.ExchangeOrder, error) {
	need, ok := repository.GetNeedByID(needID)
	if !ok {
		return model.ExchangeOrder{}, bizerr.ErrNeedNotFound
	}
	if need.Status != constants.NeedStatusOpen {
		return model.ExchangeOrder{}, bizerr.ErrNeedClosed
	}
	resp, ok := repository.GetResponseByID(in.ResponseID)
	if !ok || resp.NeedID != needID {
		return model.ExchangeOrder{}, bizerr.ErrResponseNotFound
	}
	if resp.Status != constants.ResponseStatusPending {
		return model.ExchangeOrder{}, bizerr.ErrResponseNotPending
	}
	if !slices.Contains(resp.TimeSlots, in.TimeSlot) {
		return model.ExchangeOrder{}, bizerr.ErrInvalidTimeSlot
	}
	return repository.ApplySelection(needID, in.ResponseID, in.TimeSlot)
}

// ConfirmOrder 参与者确认交换单；同时段已有确认预约时提示冲突并保留原预约
func ConfirmOrder(orderID int, user string) (model.ExchangeOrder, error) {
	user = strings.TrimSpace(user)
	if user == "" {
		return model.ExchangeOrder{}, bizerr.ErrValidation
	}
	order, ok := repository.GetOrderByID(orderID)
	if !ok {
		return model.ExchangeOrder{}, bizerr.ErrOrderNotFound
	}
	if order.Status != constants.OrderStatusPending {
		return model.ExchangeOrder{}, bizerr.ErrOrderNotPending
	}
	if !order.Involves(user) {
		return model.ExchangeOrder{}, bizerr.ErrNotParticipant
	}
	if order.ConfirmedBy(user) {
		return model.ExchangeOrder{}, bizerr.ErrAlreadyConfirmed
	}
	if repository.HasConfirmedOrderAt(user, order.TimeSlot, order.ID) {
		return model.ExchangeOrder{}, bizerr.ScheduleConflict(order.TimeSlot)
	}
	return repository.ApplyConfirmation(orderID, user)
}

func ListOrders() []model.ExchangeOrder { return repository.ListOrders() }
