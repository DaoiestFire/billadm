package operation

import (
	"github.com/spf13/cobra"

	"ljw/billadm/cmd/options"
	constant "ljw/billadm/const"
	"ljw/billadm/pkg/manager"
	"ljw/billadm/utils/logger"
)

// 基于工厂模式实现子命令的注册

type CommandRegister struct {
	data map[string]*cobra.Command
	opts *options.Options
}

func NewCommandRegister(opts *options.Options) *CommandRegister {
	cr := &CommandRegister{
		data: make(map[string]*cobra.Command, 0),
		opts: opts,
	}

	cr.register(Activate, NewActivateCommand)

	cr.register(constant.Get, NewGetCommand)
	cr.register(constant.Delete, NewDeleteCommand)
	cr.register(constant.Create, NewCreateCommand)
	cr.register(constant.Edit, NewEditCommand)

	return cr
}

func (cr *CommandRegister) register(name string, constructor func(opts *options.Options) *cobra.Command) {
	cr.data[name] = constructor(cr.opts)
}

func (cr *CommandRegister) BindToCommand(command *cobra.Command) {
	for _, c := range cr.data {
		command.AddCommand(c)
	}
}

const (
	Activate = "activate"
)

func NewActivateCommand(opts *options.Options) *cobra.Command {
	command := &cobra.Command{
		Use: Activate,
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
