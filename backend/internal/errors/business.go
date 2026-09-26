package errors

import "net/http"

// 业务错误码，前后端共用同一套语义
const (
	CodeValidation         = "VALIDATION_FAILED"
	CodeInvalidID          = "INVALID_ID"
	CodeNeedNotFound       = "NEED_NOT_FOUND"
	CodeNeedClosed         = "NEED_CLOSED"
	CodeSelfResponse       = "SELF_RESPONSE"
	CodeDuplicateResponse  = "DUPLICATE_RESPONSE"
	CodeResponseNotFound   = "RESPONSE_NOT_FOUND"
	CodeResponseNotPending = "RESPONSE_NOT_PENDING"
	CodeInvalidTimeSlot    = "INVALID_TIME_SLOT"
	CodeInvalidMode        = "INVALID_MODE"
	CodeOrderExists        = "ORDER_EXISTS"
	CodeOrderNotFound      = "ORDER_NOT_FOUND"
	CodeOrderNotPending    = "ORDER_NOT_PENDING"
	CodeNotParticipant     = "NOT_PARTICIPANT"
	CodeAlreadyConfirmed   = "ALREADY_CONFIRMED"
	CodeScheduleConflict   = "SCHEDULE_CONFLICT"
	CodeInternal           = "INTERNAL_ERROR"
)

type BusinessError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"-"`
}

func (e BusinessError) Error() string { return e.Message }

func New(code string, status int, message string) BusinessError {
	return BusinessError{Code: code, Status: status, Message: message}
}

var (
	ErrValidation         = New(CodeValidation, http.StatusBadRequest, "请完整填写必填信息")
	ErrInvalidID          = New(CodeInvalidID, http.StatusBadRequest, "无效的编号")
	ErrNeedNotFound       = New(CodeNeedNotFound, http.StatusNotFound, "需求不存在")
	ErrNeedClosed         = New(CodeNeedClosed, http.StatusConflict, "该需求已进入匹配流程，无法继续响应")
	ErrSelfResponse       = New(CodeSelfResponse, http.StatusBadRequest, "不能响应自己发布的需求")
	ErrDuplicateResponse  = New(CodeDuplicateResponse, http.StatusConflict, "你已响应过该需求，请勿重复提交")
	ErrResponseNotFound   = New(CodeResponseNotFound, http.StatusNotFound, "响应不存在或不属于该需求")
	ErrResponseNotPending = New(CodeResponseNotPending, http.StatusConflict, "该响应不在等待中，无法选择")
	ErrInvalidTimeSlot    = New(CodeInvalidTimeSlot, http.StatusBadRequest, "交换时段需从响应者提供的空闲时段中选择")
	ErrInvalidMode        = New(CodeInvalidMode, http.StatusBadRequest, "交换方式仅限线上或线下")
	ErrOrderExists        = New(CodeOrderExists, http.StatusConflict, "该需求已生成交换单，不能重复生成")
	ErrOrderNotFound      = New(CodeOrderNotFound, http.StatusNotFound, "交换单不存在")
	ErrOrderNotPending    = New(CodeOrderNotPending, http.StatusConflict, "交换单已确认，无需重复操作")
	ErrNotParticipant     = New(CodeNotParticipant, http.StatusForbidden, "只有交换双方可以确认该交换单")
	ErrAlreadyConfirmed   = New(CodeAlreadyConfirmed, http.StatusConflict, "你已确认过该交换单，不能重复确认")
	ErrInternal           = New(CodeInternal, http.StatusInternalServerError, "服务器开小差了，请稍后再试")
)

// ScheduleConflict 接受前发现同时段已有确认预约时返回，原预约保留
func ScheduleConflict(slot string) BusinessError {
	return New(CodeScheduleConflict, http.StatusConflict, "你在「"+slot+"」已有确认预约，系统已保留原预约，本次确认未生效")
}
