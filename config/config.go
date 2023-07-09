package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Commands Commands `yaml:"commands"`
	Settings Settings `yaml:"settings"`
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	config := new(Config)

	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
