package constants

// 需求状态
const (
	NeedStatusOpen     = "开放中"
	NeedStatusMatching = "匹配中"
	NeedStatusDone     = "已约成"
)

// 响应状态
const (
	ResponseStatusPending     = "等待中"
	ResponseStatusSelected    = "已选中"
	ResponseStatusNotSelected = "未选中"
)

// 交换单状态
const (
	OrderStatusPending   = "待确认"
	OrderStatusConfirmed = "已确认"
)

// 交换方式
const (
	ModeOnline  = "线上"
	ModeOffline = "线下"
)

// 已确认预约在时间线上的展示状态
const AppointmentStatusConfirmed = "双方已确认"

var ExchangeModes = []string{ModeOnline, ModeOffline}

// 统一的空闲时段词表，响应与冲突检测都基于它
var TimeSlotOptions = []string{"周一晚", "周二晚", "周三晚", "周四晚", "周五晚", "周六上午", "周六下午", "周日全天"}
