package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"
	execbeat "github.com/sonnylaskar/execbeat/beater"
)

var version = "3.3.2"
var name = "execbeat"

func main() {
	err := beat.Run(name, version, execbeat.New)
	if err != nil {
		os.Exit(1)
	}
}
