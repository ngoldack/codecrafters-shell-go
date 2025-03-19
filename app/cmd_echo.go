package main

import "fmt"

func CommandEcho() CommandFunc {
	return func(args []string) error {
		if len(args) <= 1 {
			return nil
		}

		for _, v := range args[1:] {
			fmt.Print(v)
		}
		fmt.Print("\n")
		return nil
	}
}
