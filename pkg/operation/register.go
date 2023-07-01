package operation

import (
	"github.com/spf13/cobra"
	"ljw/billadm/pkg/manager"
	"ljw/billadm/utils/logger"

	"ljw/billadm/cmd/options"
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

	cr.register(Init, NewInitCommand)
	cr.register(Activate, NewActivateCommand)

	cr.register(Get, NewGetCommand)
	cr.register(Delete, NewDeleteCommand)
	cr.register(Create, NewCreateCommand)
	cr.register(Modify, NewModifyCommand)

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
	Init     = "init"
	Activate = "activate"
)

// NewInitCommand 初始化bill.config配置文件。没有的配置设置为空串，设置配置文件的创建时间和上次修改时间
func NewInitCommand(opts *options.Options) *cobra.Command {
	command := &cobra.Command{
		Use: Init,
		Run: func(cmd *cobra.Command, args []string) {
			cm := manager.GetConfigManager()
			err := cm.Save()
			if err != nil {
				logger.Errorf("failed to init bill.config --> <%v>", err)
				return
			}
			logger.Infof("initialize bill.config success")
		},
		Args: cobra.ExactArgs(0),
	}
	return command
}

func NewActivateCommand(opts *options.Options) *cobra.Command {
	command := &cobra.Command{
		Use: Activate,
		Run: func(cmd *cobra.Command, args []string) {
			cm := manager.GetConfigManager()
			err := cm.SetCurrentBillName(args[0])
			if err != nil {
				logger.Errorf("failed to set CurrentBillName --> <%v>", err)
				return
			}
			logger.Infof("set CurrentBillName success")
		},
		Args: cobra.ExactArgs(1),
	}
	return command
}
