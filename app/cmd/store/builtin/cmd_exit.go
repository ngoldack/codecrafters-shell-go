package builtin

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/codecrafters-io/shell-starter-go/app/cmd"
	"github.com/codecrafters-io/shell-starter-go/app/state"
)

func commandExit() cmd.CommandFunc {
	return func(_ *state.State, args []string) error {
		if len(args) <= 1 {
			return errors.New("no status code")
		}

		i, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("failed to convert to int: %w", err)
		}

		os.Exit(i)
		return nil
	}
}
