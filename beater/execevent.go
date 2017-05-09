package beat

import (
	"github.com/elastic/beats/libbeat/common"
	"time"
)

type ExecEvent struct {
	ReadTime     time.Time
	DocumentType string
	Fields       map[string]string
	Exec         *Exec
	Line         *Line
}

type Exec struct {
	Command  string `json:"command,omitempty"`
	StdOut   string `json:"stdout"`
	StdErr   string `json:"stderr,omitempty"`
	ExitCode int    `json:"exitCode"`
}

type Line struct {
	Command    string `json:"command,omitempty"`
	Source     string `json:"source"`
	LineNumber int    `json:"line_number"`
	Line       string `json:"line"`
	ExitCode   int    `json:"exitCode"`
}

func (h *ExecEvent) ToMapStr() common.MapStr {
	event := common.MapStr{
		"@timestamp": common.Time(h.ReadTime),
		"type":       h.DocumentType,
	}
	if h.Exec != nil {
		event["exec"] = h.Exec
	}
	if h.Line != nil {
		event["line"] = h.Line
	}

	if h.Fields != nil {
		event["fields"] = h.Fields
	}

	return event
}
