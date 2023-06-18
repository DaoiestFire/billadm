package operation

import "github.com/spf13/cobra"

const (
	Modify = "modify"
)

func NewModifyCommand() *cobra.Command {
	command := &cobra.Command{}
	return command
}
