package builtin

import (
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/shell/cmd"
	"github.com/codecrafters-io/shell-starter-go/shell/state"
)

func commandPwd() cmd.CommandFunc {
	return func(s *state.State, args []string) error {
		fmt.Println(s.Wd)
		return nil
	}
}
