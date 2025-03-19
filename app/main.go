package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// args index 0 is always the command itself
type CommandFunc = func(args []string) error

var cmds map[string]CommandFunc = map[string]CommandFunc{
	"exit": func(args []string) error {
		if len(args) <= 1 {
			return errors.New("no status code")
		}

		i, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("failed to convert to int: %w", err)
		}

		os.Exit(i)
		return nil
	},
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err)
		}

		args := strings.Split(command, " ")

		cmd, err := getCommand(args)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		err = cmd(args)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
	}
}

func getCommand(command []string) (CommandFunc, error) {
	if len(command) == 0 {
		return nil, errors.New("empty command")
	}

	cmd, ok := cmds[command[0]]
	if !ok {
		return nil, fmt.Errorf("%s: command not found", command[:len(command)-1])
	}

	return cmd, nil
}
