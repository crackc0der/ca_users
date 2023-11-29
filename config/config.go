package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DataBase struct {
		DBHost     string `yaml:"dbHost"`
		DBPort     string `yaml:"dbPort"`
		DBName     string `yaml:"dbName"`
		DBUser     string `yaml:"dbUser"`
		DBPassword string `yaml:"dbPassword"`
	} `yaml:"dataBase"`

	Host struct {
		HostPort string `yaml:"hostPort"`
	} `yaml:"host"`
}

func NewConfig() (*Config, error) {
	var config Config

	configFile, err := os.Open("config/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("error decode config file: %w", err)
	}

	defer configFile.Close()

	configBytes, err := io.ReadAll(configFile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(configBytes, &config)

	if err != nil {
		return nil, fmt.Errorf("error unmarshal yaml config: %w", err)
	}

	return &config, nil
}
