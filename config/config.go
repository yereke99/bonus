package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	UserName       string         `yaml:"USER_NAME"`
	Password       string         `yaml:"PASSWORD"`
	DatabaseConfig DatabaseConfig `yaml:"database"` // Embed the DatabaseConfig struct
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
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
