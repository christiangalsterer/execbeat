package config

// Defaults for config variables which are not set
const (
	DefaultSchedule     string = ""
	DefaultDocumentType string = "execbeat"
)

type ExecbeatConfig struct {
	Commands []ExecConfig
}

type ExecConfig struct {
	Schedule     string
	Command      string
	Args         string
	DocumentType string            `config:"document_type"`
	Fields       map[string]string `config:"fields"`
	Name         string            `config:"name"`
}

type ConfigSettings struct {
	Execbeat ExecbeatConfig
}
