package config

import (
	"path/filepath"
	"testing"

	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	absPath, err := filepath.Abs("../tests/files/")

	assert.NotNil(t, absPath)
	assert.Nil(t, err)

	config := &ConfigSettings{}

	err = cfgfile.Read(config, absPath+"/config.yml")
	assert.Nil(t, err)

	commands := config.Execbeat.Commands
	assert.Equal(t, 2, len(commands))

	assert.Equal(t, "echo1", commands[0].Command)
	assert.Equal(t, "Execbeat", commands[0].Args)
	assert.Equal(t, "@every 1m", commands[0].Schedule)
	assert.Equal(t, 2, len(commands[0].Fields))
	assert.Equal(t, "exec", commands[0].DocumentType)

	assert.Equal(t, "echo2", commands[1].Command)
	assert.Equal(t, "Hello World", commands[1].Args)
	assert.Equal(t, "@every 2m", commands[1].Schedule)
	assert.Equal(t, 0, len(commands[1].Fields))
	assert.Equal(t, "", commands[1].DocumentType)
}
