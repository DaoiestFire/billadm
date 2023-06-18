package operation

import (
	"github.com/spf13/cobra"

	"ljw/billadm/cmd/options"
	"ljw/billadm/pkg/handler"
)

const (
	Get = "get"
)

func NewGetCommand(opts *options.Options) *cobra.Command {
	command := &cobra.Command{
		Use: Get,
		Run: func(cmd *cobra.Command, args []string) {
			handler.NewResourceHandler().Run(Get, opts)
		},
		Args: cobra.ExactArgs(1),
	}
	return command
}
