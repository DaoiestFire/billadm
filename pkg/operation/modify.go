package operation

import (
	"fmt"

	"github.com/spf13/cobra"

	"ljw/billadm/cmd/options"
	"ljw/billadm/pkg/handler"
	"ljw/billadm/utils"
)

const (
	Modify = "modify"
)

func NewModifyCommand(opts *options.Options) *cobra.Command {
	command := &cobra.Command{
		Use: Modify,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if !utils.IsResourceValid(args[0]) {
				return fmt.Errorf("resource [%s] isn't supported", args[0])
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			handler.NewResourceHandler().Run(Modify, opts)
		},
		Args: cobra.ExactArgs(1),
	}
	return command
}
