package service

import (
	"strings"

	"cyskillswap/internal/constants"
	"cyskillswap/internal/errors"
	"cyskillswap/internal/model"
	"cyskillswap/internal/repository"
)

// ListResponses 返回某条需求下的全部响应。
func ListResponses(needID int) ([]model.Response, error) {
	if _, ok := repository.GetNeed(needID); !ok {
		return nil, errors.NeedNotFound()
	}
	return repository.ListResponses(needID), nil
}

// CreateResponse 校验入参后，由响应者针对需求提交响应。
func CreateResponse(needID int, input model.ResponseInput) (model.Response, error) {
	input.Responder = strings.TrimSpace(input.Responder)
	input.OfferSkill = strings.TrimSpace(input.OfferSkill)
	input.TimeSlot = strings.TrimSpace(input.TimeSlot)
	input.PlaceType = strings.TrimSpace(input.PlaceType)
	input.Place = strings.TrimSpace(input.Place)
	input.Note = strings.TrimSpace(input.Note)

	if input.Responder == "" {
		return model.Response{}, errors.Validation("响应者不能为空")
	}
	if input.OfferSkill == "" {
		return model.Response{}, errors.Validation("请填写能提供的技能")
	}
	if !constants.IsValidTimeSlot(input.TimeSlot) {
		return model.Response{}, errors.Validation("请选择有效的空闲时段")
	}
	if !constants.IsValidPlaceType(input.PlaceType) {
		return model.Response{}, errors.Validation("请选择线上或线下地点类型")
	}
	if input.Place == "" {
		return model.Response{}, errors.Validation("请填写线上或线下地点")
	}
	return repository.CreateResponse(needID, input)
}
