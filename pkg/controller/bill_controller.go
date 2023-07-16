package controller

import (
	"ljw/billadm/cmd/options"
	v1 "ljw/billadm/pkg/api/v1"
	"ljw/billadm/pkg/storage"
)

var _ v1.Controller = &BillController{}

type BillController struct {
}

func (b *BillController) Get(storage *storage.Storage, config *options.Config) error {

}

func (b *BillController) Create(storage *storage.Storage, config *options.Config) error {

}

func (b *BillController) Delete(storage *storage.Storage, config *options.Config) error {

}

func (b *BillController) Edit(storage *storage.Storage, config *options.Config) error {

}
