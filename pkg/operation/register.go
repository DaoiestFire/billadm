package operation

import (
	"github.com/spf13/cobra"

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
	Init = "init"
)

func NewInitCommand(opts *options.Options) *cobra.Command {
	command := &cobra.Command{
		Use: Init,
		Run: func(cmd *cobra.Command, args []string) {
		},
		Args: cobra.ExactArgs(0),
	}
	return command
}
