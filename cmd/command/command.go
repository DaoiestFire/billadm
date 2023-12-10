package command

import (
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"k8s.io/klog/v2"

	"ljw/billadm/pkg/api/service"
	"ljw/billadm/pkg/api/v1"
	"ljw/billadm/pkg/storage"
	"ljw/billadm/utils/fileutils"
	"ljw/billadm/utils/logger"
)

func NewBillClientCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "billctl get|delete|create|edit|activate",
		Short: "billctl: a command executable for managing your bills",
	}

	cr := NewCommandRegister()
	cr.BindToCommand(cmd)
	return cmd
}

func NewBillServerCommand() *cobra.Command {
	billServerConfig := v1.NewBillServerConfig()
	cmd := &cobra.Command{
		Use:   "billserver",
		Short: "billserver is backend service providing storage functionality",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := billServerConfig.Validate(); err != nil {
				return err
			}

			logger.InitLogger(billServerConfig.LogFilePath)
			defer logger.Flush()

			err := runBillServer(billServerConfig)
			if err != nil {
				return err
			}
			return nil
		},
	}
	billServerConfig.ApplyTo(cmd.Flags())
	return cmd
}

func runBillServer(bsc *v1.BillServerConfig) error {
	if err := fileutils.RemoveDirectoryOrFile(bsc.SocketPath); err != nil {
		return err
	}
	lis, err := net.Listen("unix", bsc.SocketPath)
	if err != nil {
		return err
	}

	backend, err := storage.NewStorage(bsc.InstallPath)
	if err != nil {
		return err
	}
	klog.Info("get backend successfully")
	defer func() {
		err := backend.Finalizer()
		if err != nil {
			klog.Infof("backend save data failed: %v", err)
		}
	}()

	s := grpc.NewServer()
	service.RegisterStorageServiceServer(s, backend)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		klog.Info("bill server exit failed")
		return err
	}
	klog.Info("bill server exit successfully")
	return nil
}
