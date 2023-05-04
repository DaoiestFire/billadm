package operation

import "github.com/spf13/cobra"

const (
	Create = "create"
)

func NewCreateCommand() *cobra.Command {
	command := &cobra.Command{}
	return command
}
