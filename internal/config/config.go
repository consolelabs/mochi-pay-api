package config

import (
	"log"

	"github.com/spf13/viper"
)

// Loader load config from reader into Viper
type Loader interface {
	Load(*viper.Viper) (*viper.Viper, error)
}

type Config struct {
	DB DBConnection
}

type DBConnection struct {
	Host    string
	Port    string
	User    string
	Name    string
	Pass    string
	SSLMode string
}

// Read sets default to a viper instance and read user config to override these defaults
func Read() (*viper.Viper, error) {
	dcfg := viper.New()

	// read configs from .env file
	fileLoader := NewFileLoader(".env", ".")
	dcfg, err := fileLoader.Load(dcfg)
	if err != nil {
		log.Printf("Failed to load .env file: %s", err.Error())
	}

	// read configs from environment variables
	envLoader := NewENVLoader()
	dcfg, err = envLoader.Load(dcfg)
	if err != nil {
		log.Printf("Failed to load environment: %s", err.Error())
	}

	return dcfg, nil
}

func LoadConfig() *Config {
	v, err := Read()
	if err != nil {
		log.Printf("Failed to load config: %s", err.Error())
	}
	return generateConfig(v)
}

func generateConfig(v *viper.Viper) *Config {
	return &Config{
		DB: DBConnection{
			Host:    v.GetString("DB_HOST"),
			Port:    v.GetString("DB_PORT"),
			User:    v.GetString("DB_USER"),
			Name:    v.GetString("DB_NAME"),
			Pass:    v.GetString("DB_PASS"),
			SSLMode: v.GetString("DB_SSL_MODE"),
		},
	}
}
