package controller

import (
	v1 "ljw/billadm/pkg/api/v1"
	"ljw/billadm/pkg/storage"
)

type Controller interface {
	Get(storage *storage.Storage, config *v1.Config) error
	Delete(storage *storage.Storage, config *v1.Config) error
	Create(storage *storage.Storage, config *v1.Config) error
	Edit(storage *storage.Storage, config *v1.Config) error
}
