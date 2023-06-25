package operation

import (
	"fmt"

	"github.com/spf13/cobra"

	"ljw/billadm/cmd/options"
	"ljw/billadm/pkg/handler"
	"ljw/billadm/utils"
)

const (
	Get = "get"
)

func NewGetCommand(opts *options.Options) *cobra.Command {
	command := &cobra.Command{
		Use: Get,
		Run: func(cmd *cobra.Command, args []string) {
			handler.NewResourceHandler().Run(Get, args, opts)
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 || len(args) > 1 {
				return fmt.Errorf("len of args should 1 or 2")
			}
			if !utils.IsResourceValid(args[0]) {
				return fmt.Errorf("resource [%s] isn't supported", args[0])
			}
			return nil
		},
	}
	return command
}
