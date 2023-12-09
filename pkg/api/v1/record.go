package v1

import (
	"strings"
	"time"

	metav1 "ljw/billadm/pkg/api/meta/v1"
	"ljw/billadm/pkg/api/service"
)

type RecordMap map[string]IRecord

func (recordMap RecordMap) ToRecordInfoMap() map[string]*service.RecordInfo {
	recordInfoMap := make(map[string]*service.RecordInfo)
	for key := range recordMap {
		recordInfoMap[key] = recordMap[key].ToRecordInfo()
	}
	return recordInfoMap
}

func (recordMap RecordMap) FromRecordInfoMap(recordInfoMap map[string]*service.RecordInfo) {
	for key := range recordInfoMap {
		record := NewRecord("")
		record.FromRecordInfo(recordInfoMap[key])
		recordMap[key] = record
	}
}

func (recordMap RecordMap) Clone() RecordMap {
	ret := RecordMap{}
	for key := range recordMap {
		ret[key] = recordMap[key].Clone()
	}
	return ret
}

type RecordSpec struct {
	Id              string    `json:"id"`
	Cost            float32   `json:"cost"`
	Label           LabelType `json:"label"`
	Description     string    `json:"description,omitempty"`
	ConsumptionTime int64     `json:"consumptionTime,omitempty"`
}

type Record struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec RecordSpec `json:"spec,omitempty"`
}

type IRecord interface {
	GetName() string
	SetName(string)
	GetCreationTime() int64
	SetCreationTime(int64)
	GetModifyTime() int64
	SetModifyTime(int64)

	GetID() string
	SetID(string)
	GetCost() float32
	SetCost(float32)
	GetLabel() LabelType
	SetLabel(LabelType)
	GetDescription() string
	SetDescription(string)
	GetConsumptionTime() int64
	SetConsumptionTime(int64)

	ToRecordInfo() *service.RecordInfo
	FromRecordInfo(info *service.RecordInfo)

	Clone() IRecord
	GetDayEntryName() string
}

var _ IRecord = &Record{}

func NewRecord(id string) IRecord {
	record := &Record{}
	record.Kind = metav1.Record
	record.APIVersion = metav1.V1
	record.Name = id
	record.CreationTimestamp = time.Now().Unix()
	record.ModifyTimestamp = time.Now().Unix()
	record.Spec.Id = id
	return record
}

func (r *Record) GetName() string {
	return r.Name
}

func (r *Record) SetName(name string) {
	r.Name = name
}

func (r *Record) GetCreationTime() int64 {
	return r.CreationTimestamp
}

func (r *Record) SetCreationTime(t int64) {
	r.CreationTimestamp = t
}

func (r *Record) GetModifyTime() int64 {
	return r.ModifyTimestamp
}

func (r *Record) SetModifyTime(t int64) {
	r.ModifyTimestamp = t
}

func (r *Record) GetID() string {
	return r.Spec.Id
}

func (r *Record) SetID(id string) {
	r.Spec.Id = id
}

func (r *Record) GetCost() float32 {
	return r.Spec.Cost
}

func (r *Record) SetCost(cost float32) {
	r.ModifyTimestamp = time.Now().Unix()
	r.Spec.Cost = cost
}

func (r *Record) GetLabel() LabelType {
	return r.Spec.Label
}

func (r *Record) SetLabel(label LabelType) {
	r.ModifyTimestamp = time.Now().Unix()
	r.Spec.Label = label
}

func (r *Record) GetDescription() string {
	return r.Spec.Description
}

func (r *Record) SetDescription(dsp string) {
	r.ModifyTimestamp = time.Now().Unix()
	r.Spec.Description = dsp
}

func (r *Record) GetConsumptionTime() int64 {
	return r.Spec.ConsumptionTime
}

func (r *Record) SetConsumptionTime(t int64) {
	r.Spec.ConsumptionTime = t
}

func (r *Record) ToRecordInfo() *service.RecordInfo {
	return &service.RecordInfo{
		TypeMeta: &service.TypeMeta{
			Kind:       r.Kind,
			ApiVersion: r.APIVersion,
		},
		ObjectMeta: &service.ObjectMeta{
			Name:              r.Name,
			CreationTimestamp: r.CreationTimestamp,
			ModifyTimestamp:   r.ModifyTimestamp,
		},
		Id:              r.Spec.Id,
		Cost:            r.Spec.Cost,
		Label:           uint32(r.Spec.Label),
		Description:     r.Spec.Description,
		ConsumptionTime: r.Spec.ConsumptionTime,
	}
}

func (r *Record) FromRecordInfo(info *service.RecordInfo) {
	r.Kind = info.TypeMeta.Kind
	r.APIVersion = info.TypeMeta.ApiVersion
	r.Name = info.ObjectMeta.Name
	r.CreationTimestamp = info.ObjectMeta.CreationTimestamp
	r.ModifyTimestamp = info.ObjectMeta.ModifyTimestamp
	r.Spec.Id = info.Id
	r.Spec.Cost = info.Cost
	r.Spec.Label = LabelType(info.Label)
	r.Spec.Description = info.Description
	r.Spec.ConsumptionTime = info.ConsumptionTime
}

func (r *Record) Clone() IRecord {
	return &Record{
		TypeMeta: metav1.TypeMeta{
			Kind:       r.Kind,
			APIVersion: r.APIVersion,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:              r.Name,
			CreationTimestamp: r.CreationTimestamp,
			ModifyTimestamp:   r.ModifyTimestamp,
		},
		Spec: RecordSpec{
			Id:              r.Spec.Id,
			Cost:            r.Spec.Cost,
			Label:           r.Spec.Label,
			Description:     r.Spec.Description,
			ConsumptionTime: r.Spec.ConsumptionTime,
		},
	}
}

func (r *Record) GetDayEntryName() string {
	return strings.Split(r.GetName(), "/")[0]
}
