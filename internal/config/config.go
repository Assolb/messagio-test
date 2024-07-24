package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Database DatabaseConfig `yaml:"database"`
	Kafka    KafkaConfig    `yaml:"kafka"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type KafkaConfig struct {
	Broker  string `yaml:"broker"`
	Topic   string `yaml:"topic"`
	GroupId string `yaml:"groupId"`
}

var configuration *Config

func GetConfig() (*Config, error) {
	if configuration == nil {
		return nil, errors.New("config is not present")
	}
	return configuration, nil
}

func LoadConfig(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	configuration = &config

	return nil
}
