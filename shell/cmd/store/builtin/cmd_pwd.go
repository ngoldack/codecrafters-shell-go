package builtin

import (
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/app/cmd"
	"github.com/codecrafters-io/shell-starter-go/app/state"
)

func commandPwd() cmd.CommandFunc {
	return func(s *state.State, args []string) error {
		fmt.Println(s.Wd)
		return nil
	}
}
