package command

import (
	"github.com/spf13/cobra"

	"ljw/billadm/cmd/options"
)

func NewBilladmCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "billadm get|delete|create|edit|activate",
		Short: "billadm: a command executable for managing your bills",
	}

	opts := &options.Options{}
	// ApplyTo
	opts.ApplyTo(cmd.PersistentFlags())

	cr := NewCommandRegister(opts)
	cr.BindToCommand(cmd)
	return cmd
}
