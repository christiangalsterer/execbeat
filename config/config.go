// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import (
	"time"

	"github.com/elastic/beats/v7/libbeat/common"
)

// Config for this beat
type Config struct {
	Commands []*common.Config `config:"commands"`
	Period   time.Duration    `config:"period"`
	Command  string           `config:"command"`
}

// DefaultConfig for this beat
var DefaultConfig = defaultConfig()

func defaultConfig() Config {
	defaultCommandConfig := map[string]interface{}{
		"period":  "10s",
		"command": "date",
	}
	commandConfigs := []*common.Config{}
	if config, err := common.NewConfigFrom(defaultCommandConfig); err == nil {
		commandConfigs = append(commandConfigs, config)
	}

	return Config{
		Commands: commandConfigs,
	}
}
