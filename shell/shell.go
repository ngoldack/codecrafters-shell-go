package shell

import (
	"bufio"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/shell/cmd"
	"github.com/codecrafters-io/shell-starter-go/shell/cmd/store/builtin"
	"github.com/codecrafters-io/shell-starter-go/shell/cmd/store/external"
	"github.com/codecrafters-io/shell-starter-go/shell/parser"
	"github.com/codecrafters-io/shell-starter-go/shell/state"
)

type Shell struct {
	s *state.State

	reg *cmd.StoreRegister

	mode Mode
}

type Mode int

const (
	ModeInteractive Mode = iota
	ModeScript
)

func NewShell(opts ...ShellOption) *Shell {
	s := &Shell{
		s:   state.NewState(),
		reg: cmd.NewStoreRegister(),
	}

	o := append(defaultShellOptions, envShellOptions()...)
	o = append(o, opts...)
	for _, opt := range o {
		opt(s)
	}

	// Register commands (builtin and external)
	s.reg.Register(builtin.NewBuiltinStore(s.reg))
	for _, path := range strings.Split(os.Getenv("PATH"), ":") {
		s.reg.Register(external.NewExternalCommandStore(path))
	}

	return s
}

func (s *Shell) Run() {
	switch s.mode {
	case ModeInteractive:
		s.runInteractive()
	case ModeScript:
		s.runScript()
	}
}

func (s *Shell) runInteractive() {
	for {
		printPrompt()

		// Wait for user input
		command, err := readCommand()
		if err != nil {
			panic(err)
		}

		args, err := parser.Parse(command)
		if err != nil {
			printError(err)
			continue
		}

		cmd, err := cmd.GetCommand(s.reg.Stores(), args[0])
		if err != nil {
			printError(err)
			continue
		}

		err = cmd.Exec(s.s, args)
		if err != nil {
			printError(err)
		}
	}
}

func (s *Shell) runScript() {
	// TODO
}

func printPrompt() {
	print("$ ")
}

func readCommand() (string, error) {
	cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(cmd, "\n"), nil
}

func printError(err error) {
	println(err.Error())
}
