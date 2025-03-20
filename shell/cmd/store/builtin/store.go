package builtin

import (
	"github.com/codecrafters-io/shell-starter-go/shell/cmd"
)

type CommandStoreBuiltin struct {
	internalCmds map[string]cmd.CommandFunc
}

func NewBuiltinStore(register *cmd.StoreRegister) *CommandStoreBuiltin {
	s := &CommandStoreBuiltin{
		internalCmds: make(map[string]cmd.CommandFunc),
	}

	s.internalCmds["exit"] = commandExit()
	s.internalCmds["echo"] = commandEcho()
	s.internalCmds["type"] = commandType(register)
	s.internalCmds["pwd"] = commandPwd()
	s.internalCmds["cd"] = commandCd()

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
