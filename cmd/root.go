package cmd

import (
	"github.com/christiangalsterer/execbeat/beater"

	cmd "github.com/elastic/beats/v7/libbeat/cmd"
	"github.com/elastic/beats/v7/libbeat/cmd/instance"
)

// Name of this beat
var Name = "execbeat"

// Version of this beat
var Version = "5.0.0-SNAPSHOT"

// RootCmd to handle beats cli
var RootCmd = cmd.GenRootCmdWithSettings(beater.New, instance.Settings{Name: Name, Version: Version})
