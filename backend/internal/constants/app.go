package constants

const (
	ServiceName = "cyskillswap"
	APIPrefix   = "/api"
)

var SkillCategories = []string{"摄影", "编程", "乐器", "外语", "平面设计", "健身指导"}

const (
	CreditBronze = "青铜互助者"
	CreditSilver = "白银协作者"
	CreditGold   = "黄金导师"
)

// 需求状态
const (
	NeedStatusOpen   = "等待响应"
	NeedStatusBooked = "已约成"
)

// 响应状态
const (
	ResponseStatusPending  = "等待中"
	ResponseStatusChosen   = "已选中"
	ResponseStatusRejected = "未选中"
)

// 交换单状态
const (
	SwapOrderStatusPending   = "待双方确认"
	SwapOrderStatusConfirmed = "已确认"
)

// 预约状态
const (
	AppointmentStatusConfirmed = "双方已确认"
)

// 地点类型
const (
	PlaceTypeOnline  = "线上"
	PlaceTypeOffline = "线下"
)

// 可约时段，响应与交换单统一从这里取值
var TimeSlots = []string{"周二晚", "周三晚", "周六上午", "周六下午", "周日全天"}

func IsValidTimeSlot(slot string) bool {
	for _, item := range TimeSlots {
		if item == slot {
			return true
		}
	}
	return false
}

func IsValidPlaceType(placeType string) bool {
	return placeType == PlaceTypeOnline || placeType == PlaceTypeOffline
}
