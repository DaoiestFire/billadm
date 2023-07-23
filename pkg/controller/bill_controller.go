package controller

import (
	"fmt"

	logger "k8s.io/klog/v2"
	v1 "ljw/billadm/pkg/api/v1"
	"ljw/billadm/pkg/storage"
	"ljw/billadm/utils/view"
)

var _ Controller = &BillController{}

type BillController struct {
}

func (b *BillController) Get(storage *storage.Storage, config *v1.Config) error {
	bills := storage.ListAllBill()
	currentBill := storage.GetCurrentBillName()
	if err := view.PrintBills(bills, currentBill); err != nil {
		logger.Errorf("get bills failed:%v", err)
		return err
	}
	return nil
}

func (b *BillController) Create(storage *storage.Storage, config *v1.Config) error {
	if config.Name == "" {
		return fmt.Errorf("specify [name args] to create bill")
	}
	err := storage.CreateBill(config.Name)
	if err != nil {
		return err
	}
	return nil
}

func (b *BillController) Delete(storage *storage.Storage, config *v1.Config) error {
	if config.Name == "" {
		return fmt.Errorf("specify [name args] to create bill")
	}
	err := storage.DeleteBill(config.Name)
	if err != nil {
		return err
	}
	return nil
}

func (b *BillController) Edit(storage *storage.Storage, config *v1.Config) error {
	// 如果没有指定正在活动的bill就修改当前bill，如果没有激活的当前bill就报错
	billName := config.Name
	if billName == "" {
		billName = storage.GetCurrentBillName()
	}

	if billName == "" {
		return fmt.Errorf("activate current bill or specify a bill name to edit bill")
	}

	if config.User == "" {
		return fmt.Errorf("specify a user name to edit bill")
	}

	bill, err := storage.GetBill(billName)
	if err != nil {
		return err
	}
	bill.SetUser(config.User)
	return nil
}
