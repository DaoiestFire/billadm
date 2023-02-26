package handler

import "ljw/billadm/pkg/types"

var _ handler = &BillHandler{}

type BillHandler struct {
}

func (bh *BillHandler) Get(am types.IArgsMap) error {
	return nil
}

func (bh *BillHandler) Delete(am types.IArgsMap) error {
	return nil
}

func (bh *BillHandler) Create(am types.IArgsMap) error {
	return nil
}
