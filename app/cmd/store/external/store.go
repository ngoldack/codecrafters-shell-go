package external

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/codecrafters-io/shell-starter-go/app/cmd"
	"github.com/codecrafters-io/shell-starter-go/app/state"
)

type ExternalCommandStore struct {
	path string
}

func NewExternalCommandStore(path string) ExternalCommandStore {
	return ExternalCommandStore{
		path: path,
	}
}

func (s ExternalCommandStore) Find(name string) (cmd.Command, error) {
	entries, err := os.ReadDir(s.path)
	if err != nil {
		// skip if dir can't be opened
		return cmd.Command{}, err
	}

	for _, dr := range entries {
		info, err := dr.Info()
		if err != nil {
			return cmd.Command{}, err
		}

		if info.Name() == name {
			fp := fmt.Sprintf("%s/%s", s.path, name)
			return cmd.NewCommand(name, "path", fp, executeExternalCommand()), nil
		}
	}

	return cmd.Command{}, cmd.NewErrCommandNotFound(name)
}

func executeExternalCommand() cmd.CommandFunc {
	return func(_ *state.State, args []string) error {
		c := exec.Command(args[0], args[1:]...)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		if err := c.Run(); err != nil {
			return err
		}

		return nil
	}
}
