package main

import (
	"github.com/elastic/beats/libbeat/beat"
	execbeat "github.com/christiangalsterer/execbeat/beat"
)

var Version = "1.1.0-SNAPSHOT"
var Name = "execbeat"

func main() {
	beat.Run(Name, Version, execbeat.New())
}
