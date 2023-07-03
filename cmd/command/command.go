package command

import (
	"github.com/spf13/cobra"

	"ljw/billadm/cmd/options"
	"ljw/billadm/pkg/operation"
)

func NewBilladmCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "billadm get|delete|create|modify|activate|init",
		Short: "billadm: a command executable for managing your bills",
	}

	opts := &options.Options{}
	// ApplyTo
	opts.ApplyTo(cmd.PersistentFlags())

	cr := operation.NewCommandRegister(opts)
	cr.BindToCommand(cmd)
	return cmd
}
