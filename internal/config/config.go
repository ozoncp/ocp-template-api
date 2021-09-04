package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"strconv"
	"strings"
)

const configYML = "config.yml"

var cfg *Config

func GetConfigInstance() Config {
	if cfg != nil {
		return *cfg
	}
	return Config{}
}

// Database - contains all parameters database connection
type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"database"`
	SslMode  string `yaml:"sslmode"`
	Driver   string `yaml:"driver"`
}

// Grpc - contains parameter address grpc
type Grpc struct {
	Address string `yaml:"address"`
}

// Rest - contains parameter rest json connection
type Rest struct {
	Address string `yaml:"address"`
}

// Project - contains all parameters project information
type Project struct {
	Name    string `yaml:"name"`
	Author  string `yaml:"author"`
	Version string `yaml:"version"`
}

// Prometheus - contains all parameters metrics information
type Prometheus struct {
	URI  string `yaml:"uri"`
	Port string `yaml:"port"`
}

// Jaeger - contains all parameters metrics information
type Jaeger struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// Kafka - contains all parameters kafka information
type Kafka struct {
	Topic   string   `yaml:"topic"`
	Brokers []string `yaml:"brokers"`
}

// Config - contains all configuration parameters in config package
type Config struct {
	Project    Project    `yaml:"project"`
	Grpc       Grpc       `yaml:"grpc"`
	Rest       Rest       `yaml:"rest"`
	Database   Database   `yaml:"database"`
	BatchSize  int        `yaml:"batchSize"`
	Prometheus Prometheus `yaml:"prometheus"`
	Jaeger     Jaeger     `yaml:"jaeger"`
	Kafka      Kafka      `yaml:"kafka"`
}

// ReadConfigYML - read configurations from file and init instance Config
func ReadConfigYML() error {
	if cfg != nil {
		return nil
	}

	file, err := os.Open(configYML)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(&cfg); err != nil {
		return err
	}

	return nil
}
