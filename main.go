package main

import (
	"fmt"
	"os"

	"github.com/akerl/go-indefinite-article/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
