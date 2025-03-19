package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/cmd"
	"github.com/codecrafters-io/shell-starter-go/app/cmd/store/builtin"
	"github.com/codecrafters-io/shell-starter-go/app/cmd/store/external"
)

func main() {
	// Create command store
	paths := strings.Split(os.Getenv("PATH"), ":")
	stores := make([]cmd.CommandStore, 0)
	stores = append(stores, builtin.NewBuiltinStore())
	for _, path := range paths {
		stores = append(stores, external.NewExternalCommandStore(path))
	}

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err)
		}
		command, _ = strings.CutSuffix(command, "\n")

		args := strings.Split(command, " ")

		cmd, err := cmd.GetCommand(stores, args[0])
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		err = cmd.Exec(args)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
	}
}
