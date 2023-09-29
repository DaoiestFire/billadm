package v1

import (
	metav1 "ljw/billadm/pkg/api/meta/v1"
	timeutils "ljw/billadm/utils/time"
)

type RecordSpec struct {
	ID              string    `json:"id"`
	Cost            float32   `json:"cost"`
	Label           LabelType `json:"label"`
	Description     string    `json:"description,omitempty"`
	ConsumptionTime string    `json:"consumptionTime,omitempty"`
}

type Record struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec RecordSpec `json:"spec,omitempty"`
}

type IRecord interface {
	GetID() string
	GetCreationTime() string
	GetModifyTime() string
	GetConsumptionTime() string

	GetCost() float32
	SetCost(float32)
	GetDescription() string
	SetDescription(string)
	GetLabel() LabelType
	SetLabel(LabelType)
}

var _ IRecord = &Record{}

func NewRecord(id string) *Record {
	record := &Record{}
	record.Kind = metav1.Record
	record.APIVersion = metav1.V1
	record.Name = id
	record.CreationTimestamp = timeutils.GetNowTimeString()
	record.ModifyTimestamp = timeutils.GetNowTimeString()
	record.Spec.ID = id
	return record
}

func (r *Record) GetID() string {
	return r.Spec.ID
}

func (r *Record) GetCreationTime() string {
	return r.CreationTimestamp
}

func (r *Record) GetModifyTime() string {
	return r.ModifyTimestamp
}

func (r *Record) GetConsumptionTime() string {
	return r.Spec.ConsumptionTime
}

func (r *Record) GetCost() float32 {
	return r.Spec.Cost
}

func (r *Record) SetCost(cost float32) {
	r.ModifyTimestamp = timeutils.GetNowTimeString()
	r.Spec.Cost = cost
}

func (r *Record) GetDescription() string {
	return r.Spec.Description
}

func (r *Record) SetDescription(dsp string) {
	r.ModifyTimestamp = timeutils.GetNowTimeString()
	r.Spec.Description = dsp
}

func (r *Record) GetLabel() LabelType {
	return r.Spec.Label
}

func (r *Record) SetLabel(label LabelType) {
	r.ModifyTimestamp = timeutils.GetNowTimeString()
	r.Spec.Label = label
}
