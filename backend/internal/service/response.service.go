package service

import (
	"strings"

	"cyskillswap/internal/constants"
	bizerr "cyskillswap/internal/errors"
	"cyskillswap/internal/model"
	"cyskillswap/internal/repository"
)

// CreateResponse 响应者对开放中的需求发起响应
func CreateResponse(needID int, in model.CreateResponseInput) (model.Response, error) {
	if err := validateResponseInput(in); err != nil {
		return model.Response{}, err
	}
	need, ok := repository.GetNeedByID(needID)
	if !ok {
		return model.Response{}, bizerr.ErrNeedNotFound
	}
	if need.Status != constants.NeedStatusOpen {
		return model.Response{}, bizerr.ErrNeedClosed
	}
	if in.Responder == need.Requester {
		return model.Response{}, bizerr.ErrSelfResponse
	}
	if repository.HasResponseFrom(needID, in.Responder) {
		return model.Response{}, bizerr.ErrDuplicateResponse
	}

	resp := model.Response{
		NeedID:     needID,
		Responder:  strings.TrimSpace(in.Responder),
		OfferSkill: strings.TrimSpace(in.OfferSkill),
		TimeSlots:  in.TimeSlots,
		Mode:       in.Mode,
		Place:      strings.TrimSpace(in.Place),
		Note:       strings.TrimSpace(in.Note),
		Status:     constants.ResponseStatusPending,
		CreatedAt:  repository.NowString(),
	}
	return repository.AddResponse(resp), nil
}

func ListNeedResponses(needID int) ([]model.Response, error) {
	if _, ok := repository.GetNeedByID(needID); !ok {
		return nil, bizerr.ErrNeedNotFound
	}
	return repository.ListResponsesByNeed(needID), nil
}
