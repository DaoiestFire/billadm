package controller

import (
	"context"
	"ljw/billadm/pkg/api/service"

	"ljw/billadm/pkg/api/v1"
)

type Controller interface {
	Get(ctx context.Context, storageClient service.StorageServiceClient, config *v1.Config) error
	Delete(ctx context.Context, storageClient service.StorageServiceClient, config *v1.Config) error
	Create(ctx context.Context, storageClient service.StorageServiceClient, config *v1.Config) error
	Edit(ctx context.Context, storageClient service.StorageServiceClient, config *v1.Config) error
}
