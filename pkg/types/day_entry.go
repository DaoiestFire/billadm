package types

import (
	"encoding/json"
	"fmt"
	"ljw/billadm/utils/fileutils"
)

func NewDayEntry(dayTime string) *DayEntry {
	return &DayEntry{
		CurrentId: 0,
		DayTime:   dayTime,
		Records:   make(map[string]IRecord),
	}
}

func ReadOneDayEntry(path string) (*DayEntry, error) {
	data, err := fileutils.ReadFileByte(path)
	if err != nil {
		return nil, err
	}
	de := &DayEntry{}
	err = json.Unmarshal(data, de)
	if err != nil {
		return nil, err
	}
	return de, nil
}

func SaveOneDayEntry(path string, de *DayEntry) error {
	data, err := fileutils.GenerateJsonData(de)
	if err != nil {
		return err
	}

	err = fileutils.WriteFileByte(path, data)
	if err != nil {
		return err
	}
	return nil
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

func (d *DayEntry) GetRecords() []IRecord {
	recordsList := make([]IRecord, 0, d.Len())

	for _, r := range d.Records {
		recordsList = append(recordsList, r)
	}
	return recordsList
}

func (d *DayEntry) GetRecordById(id string) (IRecord, error) {
	key := d.GetKey(id)
	r, ok := d.Records[key]
	if !ok {
		return nil, fmt.Errorf("record-key(%s) isn't exsited in [%s]", key, d.DayTime)
	}
	return r, nil
}

// AddRecord 根据当前的id值创建一个新的record
func (d *DayEntry) AddRecord() {
	r := NewRecord(d.GetNextId(), d.DayTime)
	d.Records[r.GetKey()] = r
}

func (d *DayEntry) DeleteRecord(id string) error {
	key := d.GetKey(id)
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
