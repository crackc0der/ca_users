package config

import (
	"fmt"

	//nolint
	"github.com/BurntSushi/toml"
)

type Config struct {
	DataBase struct {
		DBHost     string `toml:"dbHost"`
		DBPort     string `toml:"dbPort"`
		DBName     string `toml:"dbName"`
		DBUser     string `toml:"dbUser"`
		DBPassword string `toml:"dbPassword"`
	} `toml:"DataBase"`

	Host struct {
		HostPort string `toml:"hostPort"`
	} `toml:"Host"`
}

func NewConfig() (*Config, error) {
	var config Config

	if _, err := toml.DecodeFile("config/config.toml", &config); err != nil {
		return nil, fmt.Errorf("error decode config file: %w", err)
	}

	return &config, nil
}
