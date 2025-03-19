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
				// skip if dir can't be opened
				continue
			}

			for _, e := range entries {
				if p, ok := getSearchedExecutable(e, path, args[1]); ok {
					fmt.Printf("%s is %s\n", args[1], p)
					return nil
				}
			}
		}

		return fmt.Errorf("%s: not found", args[1])
	}
}

func getSearchedExecutable(e os.DirEntry, path, exe string) (string, bool) {
	info, err := e.Info()
	if err != nil {
		return "", false
	}

	// skip if dir
	if info.IsDir() {
		return "", false
	}

	if info.Name() == exe {
		return path + "/" + info.Name(), true
	}

	return "", false
}
