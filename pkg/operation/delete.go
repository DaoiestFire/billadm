package operation

import (
	"github.com/spf13/cobra"

	"ljw/billadm/cmd/options"
	"ljw/billadm/pkg/handler"
)

const (
	Delete = "delete"
)

func NewDeleteCommand(opts *options.Options) *cobra.Command {
	command := &cobra.Command{
		Use: Delete,
		Run: func(cmd *cobra.Command, args []string) {
			handler.NewResourceHandler().Run(Delete, opts)
		},
		Args: cobra.ExactArgs(1),
	}
	return command
}
