package types

type ComposeFile struct {
	Version  string
	Services map[string]ServiceConfig
}

type ServiceConfig struct {
	Image       string
	Ports       []string
	Environment []string
	Volumes     []string
	DependsOn   []string `yaml:"depends_on"`
}
