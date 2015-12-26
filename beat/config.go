package beat

import (
	"time"
)

// Defaults for config variables which are not set
const (
	DefaultCron	string = "@every 1m"
	DefaultTimeout time.Duration = 60 * time.Second
	DefaultDocumentType string = "execbeat"
)

type ExecbeatConfig struct {
	Execs []ExecConfig
}

type ExecConfig struct {
	Cron string
	Command string
	Args string
	DocumentType string `yaml:"document_type"`
	Fields map[string]string `yaml:"fields"`
}

type ConfigSettings struct {
	Execbeat ExecbeatConfig
}
