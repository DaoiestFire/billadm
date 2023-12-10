package controller

import (
	"context"
	"fmt"

	"k8s.io/klog/v2"

	"ljw/billadm/pkg/api/service"
	"ljw/billadm/pkg/api/v1"
	"ljw/billadm/utils/view"
)

var _ Controller = &BillController{}

type BillController struct {
}

func (b *BillController) Get(ctx context.Context, storageClient service.StorageServiceClient, config *v1.Config) error {
	listAllBillResponse, err := storageClient.ListAllBill(ctx, &service.ListAllBillRequest{})
	if err != nil {
		return err
	}
	bills := make([]v1.IBill, 0, len(listAllBillResponse.BillList))
	for i := range listAllBillResponse.BillList {
		bill := v1.NewBill("")
		bill.FromBillInfo(listAllBillResponse.BillList[i])
		bills = append(bills, bill)
	}
	getCurrentBillNameResponse, err := storageClient.GetCurrentBillName(ctx, &service.GetCurrentBillNameRequest{})
	if err != nil {
		return err
	}
	currentBill := getCurrentBillNameResponse.Name
	if err := view.PrintBills(bills, currentBill); err != nil {
		klog.Errorf("get bills failed:%v", err)
		return err
	}
	return nil
}

func (b *BillController) Create(ctx context.Context, storageClient service.StorageServiceClient, config *v1.Config) error {
	if config.Name == "" {
		return fmt.Errorf("specify [name args] to create bill")
	}
	_, err := storageClient.CreateBill(ctx, &service.CreateBillRequest{Name: config.Name})
	if err != nil {
		return err
	}
	return nil
}

func (b *BillController) Delete(ctx context.Context, storageClient service.StorageServiceClient, config *v1.Config) error {
	if config.Name == "" {
		return fmt.Errorf("specify [name args] to create bill")
	}
	_, err := storageClient.DeleteBill(ctx, &service.DeleteBillRequest{Name: config.Name})
	if err != nil {
		return err
	}
	return nil
}

func (b *BillController) Edit(ctx context.Context, storageClient service.StorageServiceClient, config *v1.Config) error {
	return fmt.Errorf("not supported")
}
