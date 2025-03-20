package state

import "os"

type State struct {
	Wd   string
	Home string
}

func NewState() *State {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return &State{
		Wd:   path,
		Home: os.Getenv("HOME"),
	}
}
