package builtin

import (
	"github.com/codecrafters-io/shell-starter-go/app/cmd"
)

type CommandStoreBuiltin struct {
	internalCmds map[string]cmd.CommandFunc
}

func NewBuiltinStore() *CommandStoreBuiltin {
	s := &CommandStoreBuiltin{
		internalCmds: make(map[string]cmd.CommandFunc),
	}

	s.internalCmds["exit"] = commandExit()
	s.internalCmds["echo"] = commandEcho()
	s.internalCmds["type"] = commandType(s.internalCmds)

	return s
}

// Find implements CommandStore.
func (c CommandStoreBuiltin) Find(name string) (cmd.Command, error) {
	for cmdName, cmdFn := range c.internalCmds {
		if name == cmdName {
			return cmd.NewCommand(name, "builtin", "", cmdFn), nil
		}
	}

	return cmd.Command{}, cmd.NewErrCommandNotFound(name)
}

var _ cmd.CommandStore = (*CommandStoreBuiltin)(nil)
