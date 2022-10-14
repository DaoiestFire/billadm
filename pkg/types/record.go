package types

import (
	"fmt"

	"k8s.io/apimachinery/pkg/util/sets"
)

var _ IRecord = &Record{}

func NewRecord(id, year, month, day string) IRecord {
	return &Record{
		Id:     id,
		Year:   year,
		Month:  month,
		Day:    day,
		Labels: sets.NewString(),
	}
}

type Record struct {
	Id string `json:"id"`

	Year  string `json:"year"`
	Month string `json:"month"`
	Day   string `json:"day"`

	Cost        float32 `json:"cost"`
	Description string  `json:"description"`

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
	return r.Year
}

func (r *Record) GetMonth() string {
	return r.Month
}

func (r *Record) GetDay() string {
	return r.Day
}

func (r *Record) GetTime() string {
	return fmt.Sprintf("%s-%s-%s", r.Year, r.Month, r.Day)
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

func (r *Record) GetLabels() sets.String {
	return r.Labels
}

func (r *Record) GetKey() string {
	return fmt.Sprintf("%s-%s", r.GetTime(), r.Id)
}
