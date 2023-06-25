package operation

import (
	"fmt"

	"github.com/spf13/cobra"

	"ljw/billadm/cmd/options"
	"ljw/billadm/pkg/handler"
	"ljw/billadm/utils"
)

const (
	Create = "create"
)

func NewCreateCommand(opts *options.Options) *cobra.Command {
	command := &cobra.Command{
		Use: Create,
		Run: func(cmd *cobra.Command, args []string) {
			handler.NewResourceHandler().Run(Create, args[0], opts)
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
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
