package main

// args index 0 is always the command itself
type CommandFunc = func(args []string) error

var cmds map[string]CommandFunc

func init() {
	cmds = map[string]CommandFunc{
		"exit": CommandExit(),
		"echo": CommandEcho(),
		"type": CommandType(),
	}
}
