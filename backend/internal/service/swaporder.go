package service

import (
	"strings"

	"cyskillswap/internal/errors"
	"cyskillswap/internal/model"
	"cyskillswap/internal/repository"
)

// SwapOrders 返回全部交换单。
func SwapOrders() []model.SwapOrder { return repository.ListSwapOrders() }

// AcceptResponse 发布者从等待中的响应里选一位，生成双方待确认的交换单。
// 同时段已有确认预约或已存在交换单时会返回业务错误，不产生任何变更。
func AcceptResponse(needID, responseID int) (model.SwapOrder, error) {
	return repository.AcceptResponse(needID, responseID)
}

// ConfirmSwapOrder 交换双方分别确认交换单，双方确认后生效。
func ConfirmSwapOrder(id int, user string) (model.SwapOrder, error) {
	if strings.TrimSpace(user) == "" {
		return model.SwapOrder{}, errors.Validation("确认人不能为空")
	}
	return repository.ConfirmSwapOrder(id, user)
}
