package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	for {
		i, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err)
		}

		if i != "" {
			break
		}
	}
}
