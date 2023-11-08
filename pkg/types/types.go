package types

type ComposeFile struct {
	Version  string
	Services map[string]ServiceConfig
}

type ServiceConfig struct {
	Image          string   `yaml:"image"`
	Container_name string   `yaml:"container_name"`
	Restart        string   `yaml:"restart"`
	Ports          []string `yaml:"ports"`
	Environment    []string `yaml:"environment"`
}
