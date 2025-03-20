package shell

import "os"

type ShellOption func(*Shell)

var defaultShellOptions = []ShellOption{
	WithMode(ModeInteractive),
}

func WithMode(mode Mode) ShellOption {
	return func(s *Shell) {
		s.mode = mode
	}
}

func envShellOptions() []ShellOption {
	var opts []ShellOption

	if os.Getenv("SHELL_MODE") == "script" {
		opts = append(opts, WithMode(ModeScript))
	}

	return opts
}
