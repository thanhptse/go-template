package config

import (
	"io/ioutil"
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

// AppConfig ...
type AppConfig struct {
	Environment string `yaml:"environment"`
	ServiceName string `yaml:"service_name"`
}

// Load represent parse config from file(local env) or variable in production
func Load(filePath string) (*AppConfig, error) {
	if len(filePath) == 0 {
		filePath = os.Getenv("CONFIG-FILE")
	}

	sugar := zap.S()
	sugar.Debug("Load config ...")
	zap.S().Debugf("CONFIG_FILE=%v", filePath)

	configBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	configBytes = []byte(os.ExpandEnv(string(configBytes)))

	cfg := &AppConfig{}

	err = yaml.Unmarshal(configBytes, cfg)
	if err != nil {
		return nil, err
	}

	zap.S().Debugf("config: %+v", cfg)

	return cfg, nil
}
