package v1

import (
	"encoding/json"
	"fmt"

	constant "ljw/billadm/const"
	metav1 "ljw/billadm/pkg/api/meta/v1"
	"ljw/billadm/utils/fileutils"
	timeutils "ljw/billadm/utils/time"
)

type IRecord interface {
	GetID() string
	GetConsumptionTime() string

	GetCost() float32
	SetCost(float32)
	GetDescription() string
	SetDescription(string)
	GetLabel() LabelType
	SetLabel(LabelType)
}

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

type IBill interface {
	SetUser(user string)
	GetUser() string
}

func NewBill(name string) *Bill {
	bill := &Bill{}
	bill.Kind = metav1.Bill
	bill.APIVersion = metav1.V1
	bill.Name = name
	bill.CreationTimestamp = timeutils.GetNowTimeString()
	bill.ModifyTimestamp = timeutils.GetNowTimeString()
	bill.Spec.User = constant.DefaultUserName
	return bill
}

func (b *Bill) GetUser() string {
	return b.Spec.User
}

func (b *Bill) SetUser(user string) {
	b.ModifyTimestamp = timeutils.GetNowTimeString()
	b.Spec.User = user
}

func (b *Bill) UnmarshalFrom(data []byte) error {
	return json.Unmarshal(data, b)
}

func (b *Bill) MarshalTo() ([]byte, error) {
	return fileutils.GenerateJsonData(b)
}

type IDayEntry interface {
	AddRecord() IRecord
	DeleteRecord(string) error
	GetRecord(string) (IRecord, error)
}

func NewDayEntry(name string) *DayEntry {
	de := &DayEntry{}
	de.Kind = metav1.DayEntry
	de.APIVersion = metav1.V1
	de.Name = name
	de.CreationTimestamp = timeutils.GetNowTimeString()
	de.ModifyTimestamp = timeutils.GetNowTimeString()
	de.Spec.Records = make(map[string]*Record, 0)
	return de
}

func (d *DayEntry) UnmarshalFrom(data []byte) error {
	return json.Unmarshal(data, d)
}

func (d *DayEntry) MarshalTo() ([]byte, error) {
	return fileutils.GenerateJsonData(d)
}

func (d *DayEntry) AddRecord() IRecord {
	r := NewRecord(d.getNextID())
	return r
}

func (d *DayEntry) DeleteRecord(id string) error {
	if _, ok := d.Spec.Records[id]; !ok {
		return fmt.Errorf("record [%s] not existed", id)
	} else {
		delete(d.Spec.Records, id)
		return nil
	}
}

func (d *DayEntry) GetRecord(id string) (IRecord, error) {
	if r, ok := d.Spec.Records[id]; !ok {
		return nil, fmt.Errorf("record [%s] not existed", id)
	} else {
		return r, nil
	}
}

func (d *DayEntry) getNextID() string {
	id := fmt.Sprintf("%03d", d.Spec.CurrentId)
	d.Spec.CurrentId++
	return id
}
