package manager

import "ljw/billadm/pkg/types"

var _ manager = &LabelManager{}

type LabelManager struct {
	data types.Data
}

func (dem *LabelManager) Get() error {
	return nil
}

func (dem *LabelManager) Delete() error { 
	return nil
}

func (dem *LabelManager) Create() error {
	return nil
}

func (dem *LabelManager) Status() error {
	return nil
}

func (dem *LabelManager) Recover() error {
	return nil
}