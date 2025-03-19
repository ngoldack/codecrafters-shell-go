package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func CommandExit() CommandFunc {
	return func(args []string) error {
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
