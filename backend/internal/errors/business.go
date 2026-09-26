package errors

import (
	"fmt"
	"net/http"
)

// BusinessError 是统一业务异常，Status 为对应的 HTTP 状态码。
type BusinessError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"-"`
}

func (e BusinessError) Error() string { return e.Message }

const (
	CodeValidation         = "VALIDATION_ERROR"
	CodeNeedNotFound       = "NEED_NOT_FOUND"
	CodeNeedClosed         = "NEED_CLOSED"
	CodeSelfResponse       = "SELF_RESPONSE"
	CodeDuplicateResponse  = "DUPLICATE_RESPONSE"
	CodeResponseNotFound   = "RESPONSE_NOT_FOUND"
	CodeResponseNotPending = "RESPONSE_NOT_PENDING"
	CodeScheduleConflict   = "SCHEDULE_CONFLICT"
	CodeDuplicateOrder     = "DUPLICATE_ORDER"
	CodeOrderNotFound      = "ORDER_NOT_FOUND"
	CodeNotParticipant     = "NOT_PARTICIPANT"
	CodeAlreadyConfirmed   = "ALREADY_CONFIRMED"
)

func Validation(message string) BusinessError {
	return BusinessError{Code: CodeValidation, Message: message, Status: http.StatusBadRequest}
}

func NeedNotFound() BusinessError {
	return BusinessError{Code: CodeNeedNotFound, Message: "需求不存在", Status: http.StatusNotFound}
}

func NeedClosed() BusinessError {
	return BusinessError{Code: CodeNeedClosed, Message: "需求已约成，无法继续操作", Status: http.StatusConflict}
}

func SelfResponse() BusinessError {
	return BusinessError{Code: CodeSelfResponse, Message: "不能响应自己发布的需求", Status: http.StatusBadRequest}
}

func DuplicateResponse() BusinessError {
	return BusinessError{Code: CodeDuplicateResponse, Message: "你已响应过该需求，请勿重复提交", Status: http.StatusConflict}
}

func ResponseNotFound() BusinessError {
	return BusinessError{Code: CodeResponseNotFound, Message: "响应不存在", Status: http.StatusNotFound}
}

func ResponseNotPending() BusinessError {
	return BusinessError{Code: CodeResponseNotPending, Message: "该响应已处理，请刷新查看最新状态", Status: http.StatusConflict}
}

func ScheduleConflict(user, slot string) BusinessError {
	return BusinessError{
		Code:    CodeScheduleConflict,
		Message: fmt.Sprintf("%s 在「%s」已有确认预约，已保留原预约，请改约其他时段", user, slot),
		Status:  http.StatusConflict,
	}
}

func DuplicateOrder() BusinessError {
	return BusinessError{Code: CodeDuplicateOrder, Message: "该需求已生成交换单，不能重复生成", Status: http.StatusConflict}
}

func OrderNotFound() BusinessError {
	return BusinessError{Code: CodeOrderNotFound, Message: "交换单不存在", Status: http.StatusNotFound}
}

func NotParticipant() BusinessError {
	return BusinessError{Code: CodeNotParticipant, Message: "只有交换双方才能确认该交换单", Status: http.StatusForbidden}
}

func AlreadyConfirmed() BusinessError {
	return BusinessError{Code: CodeAlreadyConfirmed, Message: "你已确认过该交换单", Status: http.StatusConflict}
}
