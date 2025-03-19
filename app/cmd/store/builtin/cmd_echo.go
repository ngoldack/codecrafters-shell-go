package builtin

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/cmd"
	"github.com/codecrafters-io/shell-starter-go/app/state"
)

func commandEcho() cmd.CommandFunc {
	return func(_ *state.State, args []string) error {
		if len(args) <= 1 {
			return nil
		}

		fmt.Println(strings.Join(args[1:], " "))
		return nil
	}
}
