package handler

import "ljw/billadm/pkg/types"

var _ handler = &LabelHandler{}

type LabelHandler struct {
}

func (lh *LabelHandler) Get(am types.IArgsMap) error {
	return nil
}

func (lh *LabelHandler) Delete(am types.IArgsMap) error {
	return nil
}

func (lh *LabelHandler) Create(am types.IArgsMap) error {
	return nil
}
