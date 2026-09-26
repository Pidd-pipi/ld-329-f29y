package model

// Response 是响应者针对需求提交的响应
type Response struct {
	ID         int      `json:"id"`
	NeedID     int      `json:"needId"`
	Responder  string   `json:"responder"`
	OfferSkill string   `json:"offerSkill"`
	TimeSlots  []string `json:"timeSlots"`
	Mode       string   `json:"mode"`
	Place      string   `json:"place"`
	Note       string   `json:"note"`
	Status     string   `json:"status"`
	CreatedAt  string   `json:"createdAt"`
}

// CreateResponseInput 发起响应的请求体
type CreateResponseInput struct {
	Responder  string   `json:"responder" binding:"required"`
	OfferSkill string   `json:"offerSkill" binding:"required"`
	TimeSlots  []string `json:"timeSlots" binding:"required,min=1"`
	Mode       string   `json:"mode" binding:"required"`
	Place      string   `json:"place" binding:"required"`
	Note       string   `json:"note"`
}
