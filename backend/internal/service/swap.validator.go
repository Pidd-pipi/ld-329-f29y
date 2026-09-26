package service

import (
	"slices"
	"strings"

	"cyskillswap/internal/constants"
	bizerr "cyskillswap/internal/errors"
	"cyskillswap/internal/model"
)

// validateResponseInput 校验响应必填项、交换方式与时段词表
func validateResponseInput(in model.CreateResponseInput) error {
	if strings.TrimSpace(in.Responder) == "" ||
		strings.TrimSpace(in.OfferSkill) == "" ||
		strings.TrimSpace(in.Place) == "" {
		return bizerr.ErrValidation
	}
	if len(in.TimeSlots) == 0 {
		return bizerr.ErrValidation
	}
	if !slices.Contains(constants.ExchangeModes, in.Mode) {
		return bizerr.ErrInvalidMode
	}
	for _, slot := range in.TimeSlots {
		if !slices.Contains(constants.TimeSlotOptions, slot) {
			return bizerr.ErrInvalidTimeSlot
		}
	}
	return nil
}
