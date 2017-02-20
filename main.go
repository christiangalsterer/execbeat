package main

import (
	execbeat "github.com/christiangalsterer/execbeat/beater"
	"github.com/elastic/beats/libbeat/beat"
	"os"
)

var version = "3.0.1-SNAPSHOT"
var name = "execbeat"

func main() {
	err := beat.Run(name, version, execbeat.New)
	if err != nil {
		os.Exit(1)
	}
}
