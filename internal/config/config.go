package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type (
	Application struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	Server struct {
		Port int `yaml:"port"`
	}

	KafkaTopic struct {
		Alias             string `yaml:"alias"`
		Topic             string `yaml:"topic"`
		NumPartitions     int    `yaml:"num_partitions"`
		ReplicationFactor int    `yaml:"replication_factor"`
		GroupID           string `yaml:"group_id"`
	}

	Kafka struct {
		Broker string       `yaml:"broker"`
		Topics []KafkaTopic `yaml:"topics"`
	}

	Config struct {
		App    Application `yaml:"app"`
		Server Server      `yaml:"server"`
		Kafka  Kafka       `yaml:"kafka"`
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
