package main

import (
	"fmt"
)

func CommandType() CommandFunc {
	return func(args []string) error {
		if len(args) <= 1 {
			return fmt.Errorf("no args given")
		}

		if _, ok := cmds[args[0]]; ok {
			fmt.Printf("%s is a shell builtin\n", args[0])
			return nil
		}

		return fmt.Errorf("%s: not found", args[0])
	}
}
