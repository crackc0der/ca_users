package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	DataBase struct {
		DBHost     string
		DBPort     string
		DBName     string
		DBUser     string
		DBPassword string
	}

	Host struct {
		HostPort string
	}
}

func NewConfig() (*Config, error) {
	var config Config

	if _, err := toml.DecodeFile("config/config.toml", &config); err != nil {
		return nil, err
	}
	return &config, nil
}
