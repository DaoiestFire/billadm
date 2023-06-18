package operation

import (
	"github.com/spf13/cobra"

	"ljw/billadm/cmd/options"
	"ljw/billadm/pkg/handler"
)

const (
	Modify = "modify"
)

func NewModifyCommand(opts *options.Options) *cobra.Command {
	command := &cobra.Command{
		Use: Modify,
		Run: func(cmd *cobra.Command, args []string) {
			handler.NewResourceHandler().Run(Modify, opts)
		},
		Args: cobra.ExactArgs(1),
	}
	return command
}
