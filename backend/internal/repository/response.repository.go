package repository

import (
	"slices"

	"cyskillswap/internal/model"
)

func GetNeedByID(id int) (model.Need, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()
	for _, need := range db.needs {
		if need.ID == id {
			return need, true
		}
	}
	return model.Need{}, false
}

func GetResponseByID(id int) (model.Response, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()
	for _, resp := range db.responses {
		if resp.ID == id {
			resp.TimeSlots = slices.Clone(resp.TimeSlots)
			return resp, true
		}
	}
	return model.Response{}, false
}

func ListResponses() []model.Response {
	db.mu.Lock()
	defer db.mu.Unlock()
	return cloneResponses(db.responses)
}

func ListResponsesByNeed(needID int) []model.Response {
	db.mu.Lock()
	defer db.mu.Unlock()
	out := make([]model.Response, 0, len(db.responses))
	for _, resp := range db.responses {
		if resp.NeedID == needID {
			out = append(out, resp)
		}
	}
	return cloneResponses(out)
}

func HasResponseFrom(needID int, responder string) bool {
	db.mu.Lock()
	defer db.mu.Unlock()
	for _, resp := range db.responses {
		if resp.NeedID == needID && resp.Responder == responder {
			return true
		}
	}
	return false
}

// AddResponse 追加响应并同步需求的响应计数
func AddResponse(resp model.Response) model.Response {
	db.mu.Lock()
	defer db.mu.Unlock()
	resp.ID = db.nextResponseID
	db.nextResponseID++
	resp.TimeSlots = slices.Clone(resp.TimeSlots)
	db.responses = append(db.responses, resp)
	for i := range db.needs {
		if db.needs[i].ID == resp.NeedID {
			db.needs[i].Responses++
		}
	}
	return resp
}

func cloneResponses(src []model.Response) []model.Response {
	out := make([]model.Response, len(src))
	for i, resp := range src {
		resp.TimeSlots = slices.Clone(resp.TimeSlots)
		out[i] = resp
	}
	return out
}
