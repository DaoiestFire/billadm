package types

import (
	"fmt"
	"ljw/billadm/utils/logger"
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
		logger.Errorf("Record-key(%s) isn't exsited in [%s]", key, d.DayTime)
		return nil, fmt.Errorf("can't get Record by key(%s)", key)
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
		logger.Errorf("not found record in [%s] [key : %s]", d.DayTime, key)
		return fmt.Errorf("Record-key(%s) isn't exsited", key)
	}
	delete(d.Records, key)
	return nil
}

func (d *DayEntry) Len() int {
	return len(d.Records)
}
