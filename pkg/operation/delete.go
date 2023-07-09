package operation

import (
	"fmt"

	"github.com/spf13/cobra"

	"ljw/billadm/cmd/options"
	constant "ljw/billadm/const"
	"ljw/billadm/pkg/handler"
	"ljw/billadm/utils"
)

func NewDeleteCommand(opts *options.Options) *cobra.Command {
	command := &cobra.Command{
		Use: constant.Delete,
		Run: func(cmd *cobra.Command, args []string) {
			handler.NewResourceHandler().Run(constant.Delete, args, opts)
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if !utils.IsResourceValid(args[0]) {
				return fmt.Errorf("resource [%s] isn't supported", args[0])
			}
			return nil
		},
	}
	return command
}
