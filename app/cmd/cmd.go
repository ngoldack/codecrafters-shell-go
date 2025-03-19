package cmd

import "errors"

type Command struct {
	name  string
	store string
	path  string
	cmdFn CommandFunc
}

func NewCommand(name string, store string, path string, cmdFn CommandFunc) Command {
	return Command{
		name: name,
		path: path,

		store: store,

		cmdFn: cmdFn,
	}
}

func (c Command) Exec(args []string) error {
	return c.cmdFn(args)
}

type CommandStore interface {
	Find(name string) (Command, error)
}

// args index 0 is always the command itself
type CommandFunc = func(args []string) error

type ErrCommandNotFound struct {
	command string
}

func NewErrCommandNotFound(command string) ErrCommandNotFound {
	return ErrCommandNotFound{command: command}
}

func (e ErrCommandNotFound) Error() string {
	return e.command + ": command not found"
}

func GetCommand(stores []CommandStore, name string) (Command, error) {
	for _, store := range stores {
		cmd, err := store.Find(name)
		if err == nil && !errors.Is(err, ErrCommandNotFound{}) {
			return cmd, nil
		}
	}

	return Command{}, NewErrCommandNotFound(name)
}
