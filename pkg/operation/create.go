package operation

import (
	"github.com/spf13/cobra"

	"ljw/billadm/cmd/options"
	"ljw/billadm/pkg/handler"
)

const (
	Create = "create"
)

func NewCreateCommand(opts *options.Options) *cobra.Command {
	command := &cobra.Command{
		Use: Create,
		Run: func(cmd *cobra.Command, args []string) {
			handler.NewResourceHandler().Run(Create, opts)
		},
		Args: cobra.ExactArgs(1),
	}
	return command
}
