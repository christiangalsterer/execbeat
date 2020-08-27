package beat

import (
	"github.com/elastic/beats/libbeat/common"
	"time"
)

type ExecEvent struct {
	ReadTime     time.Time
	DocumentType string
	Fields       map[string]string
	Exec         Exec
}

type Exec struct {
	Command  string `json:"command,omitempty"`
	StdOut   string `json:"stdout"`
	StdErr   string `json:"stderr,omitempty"`
	ExitCode int    `json:"exitCode"`
	Args     string `json:"args,omitempty"`
	Name     string `json:"name"`
}

func (h *ExecEvent) ToMapStr() common.MapStr {
	event := common.MapStr{
		"@timestamp": common.Time(h.ReadTime),
		"type":       h.DocumentType,
		"exec":       h.Exec,
	}

	if h.Fields != nil {
		event["fields"] = h.Fields
	}

	return event
}
