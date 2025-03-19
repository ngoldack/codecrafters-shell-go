package external

import (
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/cmd"
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
	return func(args []string) error {
		fmt.Printf("Program was passed %d args (including program name).\n", len(args))

		for i, arg := range args {
			if i == 0 {
				fmt.Printf("Arg #%d (program name): %v\n", i, arg)
			} else {
				fmt.Printf("Arg #%d: %v\n", i, arg)
			}
		}

		fmt.Printf("Program Signature: %s\n", calculateProgrammSignature())

		return nil
	}
}

// 10 digit random number
func calculateProgrammSignature() string {
	r := rand.Float64()
	return fmt.Sprintf("%d", int(r*10000000000))
}
