package config

import (
	"os"
	"gopkg.in/yaml.v2"
)

type Endpoint struct {
	Name    string            `yaml:"name"`
	URL     string            `yaml:"url"`
	Method  string            `yaml:"method"`
	Headers map[string]string `yaml:"headers"`
	Body    string            `yaml:"body"`
}

func LoadConfig(filepath string) ([]Endpoint, error) {
	var endpoints []Endpoint
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &endpoints)
	return endpoints, err
}
