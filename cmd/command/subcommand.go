package command

import (
	"fmt"

	"github.com/spf13/cobra"

	"ljw/billadm/cmd/options"
	constant "ljw/billadm/const"
	"ljw/billadm/pkg/handler"
	"ljw/billadm/utils"
	"ljw/billadm/utils/logger"
)

func NewCreateCommand(opts *options.Options) *cobra.Command {
	command := &cobra.Command{
		Use: constant.Create,
		Run: func(cmd *cobra.Command, args []string) {
			handler.NewResourceHandler().Run(constant.Create, args, opts)
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("len of args not equal zero")
			}
			if !utils.IsResourceValid(args[0]) {
				return fmt.Errorf("resource [%s] isn't supported", args[0])
			}
			return nil
		},
	}
	return command
}

func NewDeleteCommand(opts *options.Options) *cobra.Command {
	command := &cobra.Command{
		Use: constant.Delete,
		Run: func(cmd *cobra.Command, args []string) {
			handler.NewResourceHandler().Run(constant.Delete, args, opts)
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if !utils.IsResourceValid(args[0]) {
				return fmt.Errorf("resource [%s] isn't supported", args[0])
			}
			return nil
		},
	}
	return command
}

func NewGetCommand(opts *options.Options) *cobra.Command {
	command := &cobra.Command{
		Use: constant.Get,
		Run: func(cmd *cobra.Command, args []string) {
			handler.NewResourceHandler().Run(constant.Get, args, opts)
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 || len(args) > 1 {
				return fmt.Errorf("len of args should 1 or 2")
			}
			if !utils.IsResourceValid(args[0]) {
				return fmt.Errorf("resource [%s] isn't supported", args[0])
			}
			return nil
		},
	}
	return command
}

func NewEditCommand(opts *options.Options) *cobra.Command {
	command := &cobra.Command{
		Use: constant.Edit,
		Run: func(cmd *cobra.Command, args []string) {
			handler.NewResourceHandler().Run(constant.Edit, args, opts)
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if !utils.IsResourceValid(args[0]) {
				return fmt.Errorf("resource [%s] isn't supported", args[0])
			}
			return nil
		},
	}
	return command
}

func NewActivateCommand(opts *options.Options) *cobra.Command {
	command := &cobra.Command{
		Use: constant.Activate,
		Run: func(cmd *cobra.Command, args []string) {
			cm := manager.GetConfigManager()
			if !cm.IsBillExist(args[0]) {
				logger.Errorf("bill [%s] doesn't existed", args[0])
			}
			err := cm.SetCurrentBillName(args[0])
			if err != nil {
				logger.Errorf("failed to set CurrentBillName --> <%v>", err)
				return
			}
			if err := cm.Save(); err != nil {
				logger.Errorf("bill.config save failed ---> <%v>", err)
			}
			logger.Infof("set CurrentBillName success")
		},
		Args: cobra.ExactArgs(1),
	}
	return command
}
