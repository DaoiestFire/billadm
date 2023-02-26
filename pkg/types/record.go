package types

import (
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"
)

var _ IRecord = &Record{}

func NewRecord(id, year, month, day string) IRecord {
	return &Record{
		Id:     id,
		Labels: sets.NewString(),
	}
}

type Record struct {
	Id          string  `json:"id"`
	Cost        float32 `json:"cost"`
	Description string  `json:"description"`

	CreationTime    string `json:"creation_time"`
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
	return strings.Split(r.CreationTime, "-")[0]
}

func (r *Record) GetMonth() string {
	return strings.Split(r.CreationTime, "-")[1]
}

func (r *Record) GetDay() string {
	return strings.Split(r.CreationTime, "-")[2]
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
