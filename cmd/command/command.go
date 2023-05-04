package command

import (
	"github.com/spf13/cobra"

	"ljw/billadm/pkg/operation"
)

func NewBilladmCommand() *cobra.Command {
	cmd := &cobra.Command{}
	cr := operation.NewCommandRegister()
	cr.BindToCommand(cmd)
	return cmd
}
