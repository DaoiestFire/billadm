package types

import (
	"fmt"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/util/sets"

	constant "ljw/billadm/const"
)

var _ IRecord = &Record{}

func NewRecord(id, consumptionTime string) IRecord {
	return &Record{
		Id:              id,
		CreationTime:    time.Now().Format(constant.TimeFormat),
		ConsumptionTime: consumptionTime,
		Labels:          sets.NewString(),
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

	Labels sets.String `json:"labels,omitempty"`
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

func (r *Record) GetYear() string {
	return strings.Split(r.GetTime(), "-")[0]
}

func (r *Record) GetMonth() string {
	return strings.Split(r.GetTime(), "-")[1]
}

func (r *Record) GetDay() string {
	return strings.Split(r.GetTime(), "-")[2]
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

func (r *Record) AddLabel(lables ...string) {
	r.Labels.Insert(lables...)
}

func (r *Record) DeleteLabel(lables ...string) {
	r.Labels.Delete(lables...)
}

func (r *Record) GetLabels() []string {
	return r.Labels.List()
}

func (r *Record) GetKey() string {
	return fmt.Sprintf("%s-%s", r.GetTime(), r.Id)
}
