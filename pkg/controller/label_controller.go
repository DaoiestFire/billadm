package controller

import (
	"context"
	"fmt"

	"ljw/billadm/pkg/api/service"
	"ljw/billadm/pkg/api/v1"
	"ljw/billadm/utils/view"
)

var _ Controller = &LabelController{}

type LabelController struct {
}

func (l *LabelController) Get(ctx context.Context, storageClient service.StorageServiceClient, config *v1.Config) error {
	err := view.PrintLabels()
	if err != nil {
		return err
	}
	return nil
}

func (l *LabelController) Create(ctx context.Context, storageClient service.StorageServiceClient, config *v1.Config) error {
	return fmt.Errorf("not supported")
}

func (l *LabelController) Delete(ctx context.Context, storageClient service.StorageServiceClient, config *v1.Config) error {
	return fmt.Errorf("not supported")
}

func (l *LabelController) Edit(ctx context.Context, storageClient service.StorageServiceClient, config *v1.Config) error {
	return fmt.Errorf("not supported")
}
