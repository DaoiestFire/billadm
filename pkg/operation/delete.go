package operation

import "github.com/spf13/cobra"

const (
	Delete = "delete"
)

func NewDeleteCommand() *cobra.Command {
	command := &cobra.Command{}
	return command
}
