package builtin

import (
	"github.com/codecrafters-io/shell-starter-go/app/cmd"
	"github.com/codecrafters-io/shell-starter-go/app/state"
)

func commandCd() cmd.CommandFunc {
	return func(s *state.State, args []string) error {
		if len(args) <= 1 {
			return nil
		}

		// check if absolute path
		if args[1][0] == '/' {
			s.Wd = args[1]
			return nil
		}

		return nil
	}
}
