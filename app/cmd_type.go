package main

import (
	"fmt"
	"os"
	"strings"
)

func CommandType() CommandFunc {
	return func(args []string) error {
		if len(args) <= 1 {
			return fmt.Errorf("no args given")
		}

		// search for builtin
		if _, ok := cmds[args[1]]; ok {
			fmt.Printf("%s is a shell builtin\n", args[1])
			return nil
		}

		// search for path executables
		paths := os.Getenv("PATH")
		for _, path := range strings.Split(paths, ":") {
			entries, err := os.ReadDir(path)
			if err != nil {
				return err
			}

			for _, e := range entries {
				if p, ok := getSearchedExecutable(e, args[1]); ok {
					fmt.Printf("%s is %s\n", args[0], p)
					return nil
				}
			}
		}

		return fmt.Errorf("%s: not found", args[1])
	}
}

func getSearchedExecutable(e os.DirEntry, exe string) (string, bool) {
	info, err := e.Info()
	if err != nil {
		return "", false
	}

	// skip if dir
	if info.IsDir() {
		return "", false
	}

	if info.Name() == exe {
		return e.Name() + info.Name(), true
	}

	return "", false
}
