package types

import (
	"fmt"
)

func NewDayEntry(dayTime string) *DayEntry {
	return &DayEntry{
		CurrentId: 0,
		DayTime:   dayTime,
		Records:   make(map[string]IRecord),
	}
}

type DayEntry struct {
	CurrentId uint32             `json:"current_id"`
	DayTime   string             `json:"day_time"`
	Records   map[string]IRecord `json:"records,omitempty"`
}

func (d *DayEntry) GetKey(id string) string {
	return fmt.Sprintf("%s-%s", d.DayTime, id)
}

func (d *DayEntry) GetNextId() string {
	res := fmt.Sprintf("%03d", d.CurrentId)
	d.CurrentId++
	return res
}

func (d *DayEntry) GetRecordByKey(id string) (IRecord, error) {
	key := d.GetKey(id)
	r, ok := d.Records[key]
	if !ok {
		return nil, fmt.Errorf("Record-key(%s) isn't exsited in [%s]", key, d.DayTime)
	}
	return r, nil
}

// 假设id不会重复，因此不检查map存在同名Record情况
func (d *DayEntry) AddRecord(r IRecord) {
	d.Records[r.GetKey()] = r
}

func (d *DayEntry) DeleteRecord(key string) error {
	_, ok := d.Records[key]
	if !ok {
		return fmt.Errorf("not found record in [%s] [key : %s]", d.DayTime, key)
	}
	delete(d.Records, key)
	return nil
}

func (d *DayEntry) Len() int {
	return len(d.Records)
}
