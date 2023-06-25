package operation

import (
	"fmt"
	"github.com/spf13/cobra"

	"ljw/billadm/cmd/options"
	"ljw/billadm/pkg/handler"
	"ljw/billadm/utils"
)

const (
	Delete = "delete"
)

func NewDeleteCommand(opts *options.Options) *cobra.Command {
	command := &cobra.Command{
		Use: Delete,
		Run: func(cmd *cobra.Command, args []string) {
			handler.NewResourceHandler().Run(Delete, args, opts)
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return fmt.Errorf("len of args not equal two")
			}
			if !utils.IsResourceValid(args[0]) {
				return fmt.Errorf("resource [%s] isn't supported", args[0])
			}
			return nil
		},
	}
	return command
}
