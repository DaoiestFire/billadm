package types

type LabelType uint32

const (
	Food LabelType = 1 << iota
	Clothing
	Daily
	Home
	Digital
	Sports
	Beauty
	MomAndKids
	Pets
	Transport
	Car
	Housing
	Travel
	Leisure
	Education
	Health
	Life
	Public
	Business
	Charity
	Assistance
	Invest
	Insurance
	RepayLoan
	Recharge
	Income
	TransferMoney
	POBO
	CashInOrOut
	Refund
	Other
)

var LabelToChinese = map[LabelType]string{
	Food:          "餐饮美食",
	Clothing:      "服饰装扮",
	Daily:         "日用百货",
	Home:          "家居家装",
	Digital:       "数码电器",
	Sports:        "运动户外",
	Beauty:        "美容美发",
	MomAndKids:    "母婴亲子",
	Pets:          "宠物",
	Transport:     "交通出行",
	Car:           "爱车养车",
	Housing:       "住房物业",
	Travel:        "酒店旅游",
	Leisure:       "文化休闲",
	Education:     "教育培训",
	Health:        "医疗健康",
	Life:          "生活服务",
	Public:        "公共服务",
	Business:      "商业服务",
	Charity:       "公益捐赠",
	Assistance:    "互助保障",
	Invest:        "投资理财",
	Insurance:     "保险",
	RepayLoan:     "信用借还",
	Recharge:      "充值缴费",
	Income:        "收入",
	TransferMoney: "转账红包",
	POBO:          "亲友代付",
	CashInOrOut:   "账户存取",
	Refund:        "退款",
	Other:         "其他",
}

var LabelList = []LabelType{
	Food,
	Clothing,
	Daily,
	Home,
	Digital,
	Sports,
	Beauty,
	MomAndKids,
	Pets,
	Transport,
	Car,
	Housing,
	Travel,
	Leisure,
	Education,
	Health,
	Life,
	Public,
	Business,
	Charity,
	Assistance,
	Invest,
	Insurance,
	RepayLoan,
	Recharge,
	Income,
	TransferMoney,
	POBO,
	CashInOrOut,
	Refund,
	Other,
}
