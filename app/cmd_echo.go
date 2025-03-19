package main

import (
	"fmt"
	"strings"
)

func CommandEcho() CommandFunc {
	return func(args []string) error {
		if len(args) <= 1 {
			return nil
		}

		fmt.Println(strings.Join(args[1:], " "))
		return nil
	}
}
