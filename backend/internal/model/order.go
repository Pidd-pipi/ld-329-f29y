package model

// ExchangeOrder 是发布者选中响应后生成的交换单，需双方确认后生效
type ExchangeOrder struct {
	ID            int      `json:"id"`
	NeedID        int      `json:"needId"`
	NeedTitle     string   `json:"needTitle"`
	ResponseID    int      `json:"responseId"`
	Requester     string   `json:"requester"`
	Responder     string   `json:"responder"`
	OfferSkill    string   `json:"offerSkill"`
	TimeSlot      string   `json:"timeSlot"`
	Mode          string   `json:"mode"`
	Place         string   `json:"place"`
	Agenda        string   `json:"agenda"`
	Status        string   `json:"status"`
	Confirmations []string `json:"confirmations"`
	CreatedAt     string   `json:"createdAt"`
}

func (o *ExchangeOrder) Involves(user string) bool {
	return o.Requester == user || o.Responder == user
}

func (o *ExchangeOrder) ConfirmedBy(user string) bool {
	for _, name := range o.Confirmations {
		if name == user {
			return true
		}
	}
	return false
}

// SelectResponseInput 发布者选择响应的请求体
type SelectResponseInput struct {
	ResponseID int    `json:"responseId" binding:"required"`
	TimeSlot   string `json:"timeSlot" binding:"required"`
}

// ConfirmOrderInput 参与者确认交换单的请求体
type ConfirmOrderInput struct {
	User string `json:"user" binding:"required"`
}
