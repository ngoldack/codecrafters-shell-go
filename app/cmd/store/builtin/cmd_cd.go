package builtin

import (
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/cmd"
	"github.com/codecrafters-io/shell-starter-go/app/state"
)

func commandCd() cmd.CommandFunc {
	return func(s *state.State, args []string) error {
		if len(args) <= 1 {
			return nil
		}

		// check if absolute path
		newPath := getNewPath(s, args[1])

		// check if path exists
		if _, err := os.ReadDir(newPath); err != nil {
			return fmt.Errorf("cd: %s: No such file or directory", args[1])
		}

		s.Wd = newPath
		return nil
	}
}

func getNewPath(s *state.State, p string) string {
	if p == "" {
		return s.Wd
	}

	// absolute path
	if p[0] == '/' {
		return p
	}

	for _, seg := range strings.Split(p, "/") {
		if p == "~" {
			continue
		}

		if seg == ".." {
			split := strings.Split(s.Wd, "/")
			s.Wd = strings.Join(split[:len(split)-1], "/")
			continue
		}

		if seg == "." {
			continue
		}

		s.Wd += "/" + seg
	}

	// remove trailing slash
	s.Wd = strings.TrimSuffix(s.Wd, "/")

	// relative path
	return s.Wd
}
