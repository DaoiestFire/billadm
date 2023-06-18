package operation

import (
	"github.com/spf13/cobra"
)

// 基于工厂模式实现子命令的注册

type CommandRegister struct {
	data map[string]*cobra.Command
}

func NewCommandRegister() *CommandRegister {
	cr := &CommandRegister{
		data: make(map[string]*cobra.Command, 0),
	}
	cr.register(Get, NewGetCommand)
	cr.register(Delete, NewDeleteCommand)
	cr.register(Create, NewCreateCommand)
	return cr
}

func (cr *CommandRegister) register(name string, constructor func() *cobra.Command) {
	cr.data[name] = constructor()
}

func (cr *CommandRegister) BindToCommand(command *cobra.Command) {
	for _, c := range cr.data {
		command.AddCommand(c)
	}
}
