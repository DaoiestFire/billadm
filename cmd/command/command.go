package command

import (
	"github.com/spf13/cobra"
)

func NewBilladmCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "billadm get|delete|create|edit|activate",
		Short: "billadm: a command executable for managing your bills",
	}

	cr := NewCommandRegister()
	cr.BindToCommand(cmd)
	return cmd
}
