package beat

import (
	"testing"
	"time"
	"github.com/elastic/beats/libbeat/common"
	"github.com/stretchr/testify/assert"
)

func TestHttpEventToMapStr(t *testing.T) {
	now := time.Now()
	fields := make(map[string]string)
	fields["field1"] = "value1"
	fields["field2"] = "value2"
	command := Exec{}
	command.Command = "foo"
	command.StdOut = "test"

	event := ExecEvent{}
	event.Fields = fields
	event.DocumentType = "test"
	event.ReadTime = now
	event.Exec = command
	mapStr := event.ToMapStr()
	_, fieldsExist := mapStr["fields"]
	assert.True(t, fieldsExist)
	_, execExist := mapStr["exec"]
	assert.True(t, execExist)
	assert.Equal(t, "test", mapStr["type"])
	assert.Equal(t, common.Time(now), mapStr["@timestamp"])
}

func TestHttpEventToMapStrWIthEmptyFields(t *testing.T) {
	event := ExecEvent{}
	mapStr := event.ToMapStr()
	_, fieldsExist := mapStr["fields"]
	assert.False(t, fieldsExist)
}
