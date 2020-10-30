package config

type AppConfig struct {
	Environment string `yaml:"environment"`
	ServiceName string `yaml:"service_name"`
}
