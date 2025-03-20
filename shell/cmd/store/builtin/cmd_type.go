package builtin

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/shell/cmd"
	"github.com/codecrafters-io/shell-starter-go/shell/state"
)

func commandType(register *cmd.StoreRegister) cmd.CommandFunc {
	return func(_ *state.State, args []string) error {
		if len(args) <= 1 {
			return fmt.Errorf("no args given")
		}

		exe := args[1]
		for _, store := range register.Stores() {
			c, err := store.Find(exe)
			if err != nil {
				return err
			}

			if c.Store() == "builtin" {
				fmt.Printf("%s is a shell builtin\n", exe)
				return nil
			}

			fmt.Printf("%s is %s\n", exe, c.Path())
			return nil
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
