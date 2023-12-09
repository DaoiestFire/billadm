package v1

import (
	"encoding/json"
	"fmt"
	"time"

	metav1 "ljw/billadm/pkg/api/meta/v1"
	"ljw/billadm/pkg/api/service"
	"ljw/billadm/utils/fileutils"
)

type DayEntrySpec struct {
	CurrentId uint32    `json:"current_id"`
	Records   RecordMap `json:"records,omitempty"`
}

type DayEntry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec DayEntrySpec `json:"spec,omitempty"`
}

type IDayEntry interface {
	AddRecord(IRecord)
	DeleteRecord(string) error
	GetRecord(string) (IRecord, error)
	ListRecords() []IRecord

	GetName() string
	GetCreationTime() int64
	GetModifyTime() int64
	GetCurrentID() uint32
	GetLen() int

	UnmarshalFrom([]byte) error
	MarshalTo() ([]byte, error)

	ToDayEntryInfo() *service.DayEntryInfo
	FromDayEntryInfo(*service.DayEntryInfo)
}

var _ IDayEntry = &DayEntry{}

func NewDayEntry(name string) *DayEntry {
	de := &DayEntry{}
	de.Kind = metav1.DayEntry
	de.APIVersion = metav1.V1
	de.Name = name
	de.CreationTimestamp = time.Now().Unix()
	de.ModifyTimestamp = time.Now().Unix()
	de.Spec.Records = RecordMap{}
	return de
}

func (d *DayEntry) UnmarshalFrom(data []byte) error {
	return json.Unmarshal(data, d)
}

func (d *DayEntry) MarshalTo() ([]byte, error) {
	return fileutils.GenerateJsonData(d)
}

// AddRecord 如果record已存在则更新，不存在则创建
func (d *DayEntry) AddRecord(record IRecord) {
	var r IRecord
	r, ok := d.Spec.Records[record.GetName()]
	if !ok {
		r = NewRecord(d.getNextName())
		d.Spec.Records[r.GetName()] = r
	}
	r.SetCost(record.GetCost())
	r.SetLabel(record.GetLabel())
	r.SetDescription(record.GetDescription())
	r.SetConsumptionTime(record.GetConsumptionTime())
	r.SetModifyTime(time.Now().Unix())
}

func (d *DayEntry) DeleteRecord(name string) error {
	if _, ok := d.Spec.Records[name]; !ok {
		return fmt.Errorf("record [%s] not existed", name)
	}
	delete(d.Spec.Records, name)
	return nil
}

// GetRecord 获取已存在record的复制
func (d *DayEntry) GetRecord(name string) (IRecord, error) {
	if r, ok := d.Spec.Records[name]; !ok {
		return nil, fmt.Errorf("record [%s] not existed", name)
	} else {
		return r.Clone(), nil
	}
}

// ListRecords 获取所有record的复制
func (d *DayEntry) ListRecords() []IRecord {
	res := make([]IRecord, 0, len(d.Spec.Records))
	for key := range d.Spec.Records {
		res = append(res, d.Spec.Records[key].Clone())
	}
	return res
}

func (d *DayEntry) GetName() string {
	return d.Name
}

func (d *DayEntry) GetCreationTime() int64 {
	return d.CreationTimestamp
}

func (d *DayEntry) GetModifyTime() int64 {
	return d.ModifyTimestamp
}

func (d *DayEntry) GetCurrentID() uint32 {
	return d.Spec.CurrentId
}

func (d *DayEntry) GetLen() int {
	return len(d.Spec.Records)
}

func (d *DayEntry) ToDayEntryInfo() *service.DayEntryInfo {
	return &service.DayEntryInfo{
		TypeMeta: &service.TypeMeta{
			Kind:       d.Kind,
			ApiVersion: d.APIVersion,
		},
		ObjectMeta: &service.ObjectMeta{
			Name:              d.Name,
			CreationTimestamp: d.CreationTimestamp,
			ModifyTimestamp:   d.ModifyTimestamp,
		},
		CurrentId: d.Spec.CurrentId,
		RecordMap: d.Spec.Records.ToRecordInfoMap(),
	}
}

func (d *DayEntry) FromDayEntryInfo(info *service.DayEntryInfo) {
	d.Kind = info.TypeMeta.Kind
	d.APIVersion = info.TypeMeta.ApiVersion
	d.Name = info.ObjectMeta.Name
	d.CreationTimestamp = info.ObjectMeta.CreationTimestamp
	d.ModifyTimestamp = info.ObjectMeta.ModifyTimestamp
	d.Spec.CurrentId = info.CurrentId
	recordMap := RecordMap{}
	recordMap.FromRecordInfoMap(info.RecordMap)
	d.Spec.Records = recordMap
}

func (d *DayEntry) getNextName() string {
	id := fmt.Sprintf("%s/%03d", d.GetName(), d.Spec.CurrentId)
	d.Spec.CurrentId++
	return id
}
