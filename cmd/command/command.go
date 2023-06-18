package command

import (
	"github.com/spf13/cobra"

	"ljw/billadm/pkg/operation"
)

func NewBilladmCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "billadm get|delete|create|modify|activate|init",
		Short: "billadm: a command executable for managing your bills",
	}
	cr := operation.NewCommandRegister()
	cr.BindToCommand(cmd)
	return cmd
}
