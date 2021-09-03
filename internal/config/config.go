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

// NewEnv returns a new Config struct from .env
// Ability to get settings dynamically
func NewEnv() *Config {
	return &Config{
		Project: Project{
			Name:    getEnv("PROJECT_NAME", GetConfigInstance().Project.Name),
			Author:  getEnv("PROJECT_AUTHOR", GetConfigInstance().Project.Author),
			Version: getEnv("PROJECT_VERSION", GetConfigInstance().Project.Version),
		},
		Grpc: Grpc{
			Address: getEnv("GRPC_ADDRESS", GetConfigInstance().Grpc.Address),
		},
		Rest: Rest{
			Address: getEnv("REST_ADDRESS", GetConfigInstance().Rest.Address),
		},
		Database: Database{
			Host:     getEnv("DATABASE_HOST", GetConfigInstance().Database.Host),
			Port:     getEnv("DATABASE_PORT", GetConfigInstance().Database.Port),
			User:     getEnv("DATABASE_USER", GetConfigInstance().Database.User),
			Password: getEnv("DATABASE_PASSWORD", GetConfigInstance().Database.Password),
			Name:     getEnv("DATABASE_NAME", GetConfigInstance().Database.Name),
			SslMode:  getEnv("DATABASE_SSL_MODE", GetConfigInstance().Database.SslMode),
			Driver:   getEnv("DATABASE_DRIVER", GetConfigInstance().Database.Driver),
		},
		BatchSize: getEnvAsInt("BATCH_SIZE", GetConfigInstance().BatchSize),
		Prometheus: Prometheus{
			URI:  getEnv("PROMETHEUS_URI", GetConfigInstance().Prometheus.URI),
			Port: getEnv("PROMETHEUS_PORT", GetConfigInstance().Prometheus.Port),
		},
		Jaeger: Jaeger{
			Host: getEnv("JAEGER_HOST", GetConfigInstance().Jaeger.Host),
			Port: getEnv("JAEGER_PORT", GetConfigInstance().Jaeger.Port),
		},
		Kafka: Kafka{
			Topic:   getEnv("KAFKA_TOPIC", GetConfigInstance().Kafka.Topic),
			Brokers: getEnvAsSlice("KAFKA_BROKERS", GetConfigInstance().Kafka.Brokers, ","),
		},
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
