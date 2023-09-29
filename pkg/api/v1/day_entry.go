package v1

import (
	"encoding/json"
	"fmt"

	metav1 "ljw/billadm/pkg/api/meta/v1"
	"ljw/billadm/utils/fileutils"
	timeutils "ljw/billadm/utils/time"
)

type DayEntrySpec struct {
	CurrentId uint32             `json:"current_id"`
	Records   map[string]*Record `json:"records,omitempty"`
}

type DayEntry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec DayEntrySpec `json:"spec,omitempty"`
}

type IDayEntry interface {
	AddRecord() IRecord
	DeleteRecord(string) error
	GetRecord(string) (IRecord, error)
	ListRecords() []IRecord

	GetName() string
	GetCreationTime() string
	GetModifyTime() string
	GetCurrentID() uint32
	GetLen() int
}

var _ IDayEntry = &DayEntry{}

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

func (d *DayEntry) ListRecords() []IRecord {
	res := make([]IRecord, 0, len(d.Spec.Records))
	for key := range d.Spec.Records {
		res = append(res, d.Spec.Records[key])
	}
	return res
}

func (d *DayEntry) GetName() string {
	return d.Name
}

func (d *DayEntry) GetCreationTime() string {
	return d.CreationTimestamp
}

func (d *DayEntry) GetModifyTime() string {
	return d.ModifyTimestamp
}

func (d *DayEntry) GetCurrentID() uint32 {
	return d.Spec.CurrentId
}

func (d *DayEntry) GetLen() int {
	return len(d.Spec.Records)
}

func (d *DayEntry) getNextID() string {
	id := fmt.Sprintf("%03d", d.Spec.CurrentId)
	d.Spec.CurrentId++
	return id
}
