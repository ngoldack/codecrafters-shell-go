package builtin

import (
	"fmt"
	"os"
	"strconv"

	"github.com/codecrafters-io/shell-starter-go/shell/cmd"
	"github.com/codecrafters-io/shell-starter-go/shell/state"
)

func commandExit() cmd.CommandFunc {
	return func(_ *state.State, args []string) error {
		if len(args) <= 1 {
			os.Exit(0)
		}

		i, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("failed to convert to int: %w", err)
		}

		os.Exit(i)
		return nil
	}
}
