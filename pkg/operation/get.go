package operation

import "github.com/spf13/cobra"

const (
	Get = "get"
)

func NewGetCommand() *cobra.Command {
	command := &cobra.Command{}
	return command
}
