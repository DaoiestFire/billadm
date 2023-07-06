package operation

import (
	"fmt"
	constant "ljw/billadm/const"

	"github.com/spf13/cobra"

	"ljw/billadm/cmd/options"
	"ljw/billadm/pkg/handler"
	"ljw/billadm/utils"
)

func NewCreateCommand(opts *options.Options) *cobra.Command {
	command := &cobra.Command{
		Use: constant.Create,
		Run: func(cmd *cobra.Command, args []string) {
			handler.NewResourceHandler().Run(constant.Create, args, opts)
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return fmt.Errorf("len of args not equal zero")
			}
			if !utils.IsResourceValid(args[0]) {
				return fmt.Errorf("resource [%s] isn't supported", args[0])
			}
			return nil
		},
	}
	return command
}
