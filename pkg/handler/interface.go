package handler

import "ljw/billadm/pkg/types"

type handler interface {
	Get(types.IArgsMap) error
	Delete(types.IArgsMap) error
	Create(types.IArgsMap) error
}
