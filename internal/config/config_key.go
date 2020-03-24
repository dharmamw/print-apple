package config

type (
	// Config ...
	Config struct {
		Server   ServerConfig   `yaml:"server"`
		Firebase FirebaseConfig `yaml:"firebase"`
	}

	// ServerConfig ...
	ServerConfig struct {
		Port string `yaml:"port"`
	}

	// FirebaseConfig ...
	FirebaseConfig struct {
		ProjectID string `yaml:"ProjectID"`
	}
)
