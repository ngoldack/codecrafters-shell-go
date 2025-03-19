package builtin

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/cmd"
)

func commandEcho() cmd.CommandFunc {
	return func(args []string) error {
		if len(args) <= 1 {
			return nil
		}

		fmt.Println(strings.Join(args[1:], " "))
		return nil
	}
}
