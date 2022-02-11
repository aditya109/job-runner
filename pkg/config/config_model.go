package config

// Config is central container for configuration object
type Config struct {
	Server    Server   `yaml:"server"`
	Scenarios []string `yaml:"scenarios"`
}

// Server is the container for server-related configuration
type Server struct {
	Env string `yaml:"environment"`
}
