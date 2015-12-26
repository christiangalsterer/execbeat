package main

import (
	"github.com/elastic/beats/libbeat/beat"
	execbeat "github.com/christiangalsterer/execbeat/beat"
)

var Version = "1.0.0-beta1"
var Name = "execbeat"

func main() {
	beat.Run(Name, Version, execbeat.New())
}
