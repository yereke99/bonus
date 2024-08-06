package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	UserName string `yaml:"USER_NAME"`
	Password string `yaml:"PASSWORD"`
}

func NewConfig(fileName string) (*Config, error) {

	cfg := new(Config)
	if err := loadYAMLFile(fileName, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func loadYAMLFile(fileName string, cfg *Config) error {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return err
	}

	return nil
}
