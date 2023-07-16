package controller

import (
	"ljw/billadm/cmd/options"
	v1 "ljw/billadm/pkg/api/v1"
	"ljw/billadm/pkg/storage"
)

var _ v1.Controller = &RecordController{}

type RecordController struct {
}

func (r *RecordController) Get(storage *storage.Storage, config *options.Config) error {

}

func (r *RecordController) Create(storage *storage.Storage, config *options.Config) error {

}

func (r *RecordController) Delete(storage *storage.Storage, config *options.Config) error {

}

func (r *RecordController) Edit(storage *storage.Storage, config *options.Config) error {

}
