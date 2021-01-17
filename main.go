package main

import (
	"os"

	"github.com/christiangalsterer/execbeat/cmd"

	_ "github.com/christiangalsterer/execbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
