package beat

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

	execs := config.Execbeat.Execs
	assert.Equal(t, 2, len(execs))

	assert.Equal(t, "echo1", execs[0].Command)
	assert.Equal(t, "Execbeat", execs[0].Args)
	assert.Equal(t, "@every 1m", execs[0].Cron)
	assert.Equal(t, 2, len(execs[0].Fields))
	assert.Equal(t, "exec", execs[0].DocumentType)

	assert.Equal(t, "echo2", execs[1].Command)
	assert.Equal(t, "Hello World", execs[1].Args)
	assert.Equal(t, "@every 2m", execs[1].Cron)
	assert.Equal(t, 0, len(execs[1].Fields))
	assert.Equal(t, "", execs[1].DocumentType)
}

