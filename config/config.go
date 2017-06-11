package config

// Defaults for config variables which are not set
const (
	DefaultSchedule     string = "@every 1m"
	DefaultDocumentType string = "execbeat"
)

type ExecbeatConfig struct {
	Commands []ExecConfig
}

type ExecConfig struct {
	Schedule     string	`yaml:"schedule"`
	Command      string	`yaml:"command"`
	Args         string	`yaml:"args"`
	DocumentType string            `config:"document_type" yaml:"document_type`
	Fields       map[string]string `config:"fields" yaml:"fields"`
}

type ConfigSettings struct {
	Execbeat ExecbeatConfig
}
