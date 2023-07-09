package command

import (
	"github.com/spf13/cobra"

	"ljw/billadm/cmd/options"
	constant "ljw/billadm/const"
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

	cr.register(constant.Activate, NewActivateCommand)

	cr.register(constant.Get, NewGetCommand)
	cr.register(constant.Create, NewCreateCommand)
	cr.register(constant.Delete, NewDeleteCommand)
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
