package types

import (
	"fmt"

	timeutils "ljw/billadm/utils/time"
)

var _ IRecord = &Record{}

func NewRecord(id, consumptionTime string) *Record {
	return &Record{
		Id:              id,
		CreationTime:    timeutils.GetNowTimeString(),
		ConsumptionTime: consumptionTime,
	}
}

type Record struct {
	Id          string  `json:"id"`
	Cost        float32 `json:"cost"`
	Description string  `json:"description"`

	// 2006-01-02-15:04:05
	CreationTime string `json:"creation_time"`
	// 2006-01-02
	ConsumptionTime string `json:"consumption_time"`

	Label LabelType `json:"label"`
}

func (r *Record) GetID() string {
	return r.Id
}

func (r *Record) GetCost() float32 {
	return r.Cost
}

func (r *Record) GetDescription() string {
	return r.Description
}

func (r *Record) GetTime() string {
	return r.ConsumptionTime
}

func (r *Record) SetCost(cost float32) {
	r.Cost = cost
}

func (r *Record) SetDescription(description string) {
	r.Description = description
}

func (r *Record) SetLabels(label LabelType) {
	r.Label = label
}

func (r *Record) GetLabels() LabelType {
	return r.Label
}

func (r *Record) GetKey() string {
	return fmt.Sprintf("%s-%s", r.GetTime(), r.Id)
}
