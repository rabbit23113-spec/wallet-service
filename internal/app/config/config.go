package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port     string   `yaml:"port"`
	Dbconfig DBConfig `yaml:"dbconfig"`
}

type DBConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int16  `yaml:"port"`
	Database string `yaml:"database"`
	Dsn      string `yaml:"dsn"`
}

func NewConfig() *Config {
	var data *Config

	bytes, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(bytes, &data)
	if err != nil {
		panic(err)
	}

	return data
}
