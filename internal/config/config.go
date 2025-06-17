package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type (
	Application struct {
		Name    string
		Version string
	}

	Server struct {
		Port int
	}

	Kafka struct {
		Brokers []string
		Topic   string
		GroupID string
	}

	Config struct {
		App    Application
		Server Server
		Kafka  Kafka
	}
)

func NewConfig() *Config {
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	return &cfg
}
