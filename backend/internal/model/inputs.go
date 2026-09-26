package model

// ResponseInput 是发起响应的请求体。
type ResponseInput struct {
	Responder  string `json:"responder"`
	OfferSkill string `json:"offerSkill"`
	TimeSlot   string `json:"timeSlot"`
	PlaceType  string `json:"placeType"`
	Place      string `json:"place"`
	Note       string `json:"note"`
}

// ConfirmInput 是确认交换单的请求体。
type ConfirmInput struct {
	User string `json:"user"`
}
