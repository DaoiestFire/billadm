package types

import (
	"fmt"

	"k8s.io/klog/v2"
)

func NewDayEntry(year, month, day string) *DayEntry {
	return &DayEntry{
		Year:    year,
		Month:   month,
		Day:     day,
		Records: make(map[string]IRecord),
	}
}

type DayEntry struct {
	CurrentId uint32 `json:"current_id"`

	Year  string `json:"year"`
	Month string `json:"month"`
	Day   string `json:"day"`

	Records map[string]IRecord `json:"records,omitempty"`
}

func (d *DayEntry) GetKey(id string) string {
	return fmt.Sprintf("%s-%s-%s-%s", d.Year, d.Month, d.Day, id)
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
		klog.Errorf("Record-key(%s) isn't exsited", key)
		return nil, fmt.Errorf("can't get Record by id(%s)", id)
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
		klog.Errorf("Record-key(%s) isn't exsited", key)
		return fmt.Errorf("Record-key(%s) isn't exsited", key)
	}
	delete(d.Records, key)
	return nil
}

func (d *DayEntry) Len() int {
	return len(d.Records)
}
