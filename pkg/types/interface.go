package types

import "k8s.io/apimachinery/pkg/util/sets"

type IRecord interface {
	GetID() string
	GetCost() float32
	GetDescription() string

	GetYear() string
	GetMonth() string
	GetDay() string
	GetTime() string

	// 时间与id是无法被更改的，这个成员是自动生成的
	// 想要更改时间或id只能删除这个记录，增加一个新记录
	SetCost(float32)
	SetDescription(string)

	AddLabel(...string)
	DeleteLabel(...string)
	GetLabels() sets.String

	GetKey() string
}