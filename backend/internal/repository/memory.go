package repository

import (
	"cyskillswap/internal/constants"
	"cyskillswap/internal/model"
)

// 本文件只提供种子数据，运行期状态由 store.go 中的 Store 维护。

func seedSkills() []model.Skill {
	return []model.Skill{
		{ID: 1, Owner: "林澈", Title: "毕业照人像摄影", Category: "摄影", Level: 92, Campus: "东校区", Description: "提供构图、修图和毕业季跟拍，可交换吉他入门课。", TimeSlots: []string{"周三晚", "周六上午"}, Rewards: []string{"技能交换", "请吃饭"}, Portfolio: "12组校园人像作品"},
		{ID: 2, Owner: "周芮", Title: "Python 数据分析", Category: "编程", Level: 88, Campus: "中心校区", Description: "pandas、可视化、论文数据清洗辅导，接受小额报酬。", TimeSlots: []string{"周二晚", "周日全天"}, Rewards: []string{"技能交换", "小额报酬"}, Portfolio: "3份课程项目证书"},
		{ID: 3, Owner: "孟野", Title: "民谣吉他陪练", Category: "乐器", Level: 81, Campus: "西校区", Description: "节奏型、弹唱和舞台经验分享，想找人拍宣传照。", TimeSlots: []string{"周三晚", "周六上午"}, Rewards: []string{"技能交换", "无偿"}, Portfolio: "校园音乐节演出视频"},
	}
}

func seedNeeds() []model.Need {
	return []model.Need{
		{ID: 1, Requester: "孟野", Title: "找人帮忙拍乐队宣传照", Category: "摄影", Campus: "西校区", ExpectTime: "本周六上午", BudgetType: "技能交换", Description: "可交换 3 次吉他课，希望会调色和室外构图。", Status: constants.NeedStatusOpen},
		{ID: 2, Requester: "许安", Title: "求教 Python 数据分析", Category: "编程", Campus: "中心校区", ExpectTime: "周二晚", BudgetType: "小额报酬", Description: "论文问卷数据需要清洗和画图，最好有 pandas 经验。", Status: constants.NeedStatusOpen},
		{ID: 3, Requester: "林澈", Title: "想学吉他扫弦入门", Category: "乐器", Campus: "东校区", ExpectTime: "周三晚", BudgetType: "技能交换", Description: "用摄影课交换吉他基础，希望同校区或线上。", Status: constants.NeedStatusOpen},
	}
}

func seedResponses() []model.Response {
	return []model.Response{
		{ID: 1, NeedID: 1, Responder: "林澈", OfferSkill: "毕业照人像摄影", TimeSlot: "周六上午", PlaceType: constants.PlaceTypeOffline, Place: "西校区操场看台", Note: "可带两套设备，拍完当天出预览。", Status: constants.ResponseStatusPending},
		{ID: 2, NeedID: 1, Responder: "周芮", OfferSkill: "活动跟拍与调色", TimeSlot: "周六下午", PlaceType: constants.PlaceTypeOffline, Place: "西校区琴房门口", Note: "想交换吉他入门课。", Status: constants.ResponseStatusPending},
		{ID: 3, NeedID: 2, Responder: "周芮", OfferSkill: "Python 数据分析", TimeSlot: "周二晚", PlaceType: constants.PlaceTypeOnline, Place: "腾讯会议 381-204-776", Note: "可远程共享屏幕演示 pandas 清洗流程。", Status: constants.ResponseStatusPending},
		{ID: 4, NeedID: 3, Responder: "孟野", OfferSkill: "民谣吉他陪练", TimeSlot: "周三晚", PlaceType: constants.PlaceTypeOffline, Place: "东校区琴房 203", Note: "从扫弦节奏型开始带，课后发练习谱。", Status: constants.ResponseStatusPending},
	}
}

func seedMatches() []model.Match {
	return []model.Match{
		{ID: 1, Provider: "林澈", Learner: "孟野", OfferSkill: "毕业照人像摄影", WantedSkill: "民谣吉他陪练", Score: 96, CommonSlots: []string{"周三晚", "周六上午"}, Recommendation: "互补技能明确，双方均接受技能交换。"},
		{ID: 2, Provider: "周芮", Learner: "许安", OfferSkill: "Python 数据分析", WantedSkill: "论文数据清洗", Score: 89, CommonSlots: []string{"周二晚"}, Recommendation: "时间匹配且需求描述命中 pandas/可视化。"},
		{ID: 3, Provider: "孟野", Learner: "林澈", OfferSkill: "民谣吉他陪练", WantedSkill: "宣传照拍摄", Score: 91, CommonSlots: []string{"周六上午"}, Recommendation: "互换回报类型一致，信用分权重较高。"},
	}
}

func seedAppointments() []model.Appointment {
	return []model.Appointment{
		{ID: 1, Pair: "林澈 ↔ 孟野", Time: "周六 10:00", TimeSlot: "周六上午", Place: "东校区湖边", Status: constants.AppointmentStatusConfirmed, Agenda: "先拍宣传照，再约 2 次吉他课", Participants: []string{"林澈", "孟野"}},
		{ID: 2, Pair: "周芮 ↔ 许安", Time: "周二 19:30", TimeSlot: "周二晚", Place: "线上会议室", Status: "等待对方确认", Agenda: "导入问卷 CSV 并完成基础可视化", Participants: []string{"周芮", "许安"}},
	}
}

func seedReviews() []model.Review {
	return []model.Review{
		{ID: 1, From: "孟野", To: "林澈", Rating: 5, Content: "构图建议很细，成片当天就给了预览。"},
		{ID: 2, From: "林澈", To: "孟野", Rating: 5, Content: "吉他入门节奏拆得很清楚，课后还发了练习谱。"},
	}
}

func seedMessages() []model.Conversation {
	return []model.Conversation{
		{ID: 1, WithUser: "孟野", Unread: 2, Messages: []string{"周六湖边光线不错", "我带两套衣服可以吗？"}},
		{ID: 2, WithUser: "系统通知", Unread: 1, Messages: []string{"你与周芮的 Python 数据分析预约待确认。"}},
	}
}

func baseProfile() model.Profile {
	return model.Profile{
		Name: "林澈", Major: "新闻传播 2023", CreditScore: 91, CreditLevel: constants.CreditGold,
		Radar:   map[string]int{"摄影": 92, "修图": 86, "沟通": 90, "编程": 42, "乐器": 35},
		History: []string{"完成毕业照拍摄交换", "响应 Python 数据分析需求", "预约吉他入门课"},
	}
}
