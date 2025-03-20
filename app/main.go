package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/cmd"
	"github.com/codecrafters-io/shell-starter-go/app/cmd/store/builtin"
	"github.com/codecrafters-io/shell-starter-go/app/cmd/store/external"
	"github.com/codecrafters-io/shell-starter-go/app/state"
)

type State struct {
	Stores []cmd.CommandStore
	Wd     string
}

func main() {
	// Create command store
	paths := strings.Split(os.Getenv("PATH"), ":")
	stores := make([]cmd.CommandStore, 0)
	stores = append(stores, builtin.NewBuiltinStore())
	for _, path := range paths {
		stores = append(stores, external.NewExternalCommandStore(path))
	}

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	s := &state.State{
		Wd:   path,
		Home: os.Getenv("HOME"),
	}

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err)
		}
		command, _ = strings.CutSuffix(command, "\n")

		args, err := getArgs(command)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		cmd, err := cmd.GetCommand(stores, args[0])
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		err = cmd.Exec(s, args)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
	}
}

const singleQuote = '\''
const doubleQuote = '"'
const escapeChar = '\\'
const spaceChar = ' '

func getArgs(command string) ([]string, error) {
	args := make([]string, 0)
	var currentArg strings.Builder
	inQuotes := false
	quoteChar := rune(0)
	escapedAt := -1

	for i := 0; i < len(command); i++ {
		c := rune(command[i])

		switch {
		case c == singleQuote || c == doubleQuote:
			if escapedAt == i-1 {
				currentArg.WriteRune(c)
				escapedAt = -1
			} else if !inQuotes {
				inQuotes = true
				quoteChar = c
			} else if c == quoteChar {
				inQuotes = false
				quoteChar = 0
			} else {
				currentArg.WriteRune(c)
			}
		case c == spaceChar:
			if escapedAt == i-1 {
				currentArg.WriteRune(c)
				escapedAt = -1
			} else if inQuotes {
				currentArg.WriteRune(c)
			} else if currentArg.Len() > 0 {
				args = append(args, currentArg.String())
				currentArg.Reset()
			}
		case c == escapeChar:
			if inQuotes {
				currentArg.WriteRune(c)
			} else if escapedAt == i-1 {
				currentArg.WriteRune(c)
				escapedAt = -1
			} else {
				escapedAt = i
			}
		default:
			currentArg.WriteRune(c)
		}
	}

	if currentArg.Len() > 0 {
		args = append(args, currentArg.String())
	}

	return args, nil
}
