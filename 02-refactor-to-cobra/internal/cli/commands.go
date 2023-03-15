package cli

import (
	"github.com/spf13/cobra"
)

// CobraFn function definition of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const idFlag = "id"

func NewCommand(commandName string) *cobra.Command {
	switch commandName {
	case "beers":
		return InitBeersCmd()
	case "stores":
		return InitStoreCmd()
	}

	return nil
}
