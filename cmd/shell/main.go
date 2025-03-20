package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/shell/cmd"
	"github.com/codecrafters-io/shell-starter-go/shell/cmd/store/builtin"
	"github.com/codecrafters-io/shell-starter-go/shell/cmd/store/external"
	"github.com/codecrafters-io/shell-starter-go/shell/state"
)

func main() {
	// Create command store
	paths := strings.Split(os.Getenv("PATH"), ":")

	reg := cmd.NewStoreRegister()

	reg.Register(builtin.NewBuiltinStore(reg))
	for _, path := range paths {
		reg.Register(external.NewExternalCommandStore(path))
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

		cmd, err := cmd.GetCommand(reg.Stores(), args[0])
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

var escapedCharsInDoubleQuotes = []rune{
	'n',
	'\\',
	'$',
	'"',
}

func getArgs(command string) ([]string, error) {
	args := make([]string, 0)
	var currentArg strings.Builder
	inQuotes := false
	quoteChar := rune(0)
	escaped := false

	for i := 0; i < len(command); i++ {
		c := rune(command[i])

		if escaped {
			currentArg.WriteRune(c)
			escaped = false
			continue
		}

		switch c {
		case escapeChar:
			if inQuotes && quoteChar == singleQuote {
				currentArg.WriteRune(c)
			} else {
				escaped = true
			}
		case singleQuote, doubleQuote:
			if !inQuotes {
				inQuotes = true
				quoteChar = c
			} else if c == quoteChar {
				inQuotes = false
				quoteChar = 0
			} else {
				currentArg.WriteRune(c)
			}
		case spaceChar:
			if inQuotes {
				currentArg.WriteRune(c)
			} else if currentArg.Len() > 0 {
				args = append(args, currentArg.String())
				currentArg.Reset()
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
