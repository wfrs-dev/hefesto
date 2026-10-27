package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	SSLMode  string `json:"sslmode"`
}

type Server struct {
	Listen string `json:"listen"`
	Port     int    `json:"port"`
	LogLevel string `json:"logLevel"`
}

type Config struct {
	Database Database `json:"database"`
	Server   Server   `json:"server"`
}

func Load(path string) (*Config, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file `%s` not found", path)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file `%s`: %w", path, err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func GenerateDefaultConfig() error {
	sample := Config{
		Database: Database{
			Host:     "127.0.0.1",
			Port:     5432,
			User:     "db_user",
			Password: "db_password",
			Database: "db_name",
			SSLMode:  "disable",
		},
		Server: Server{
			Listen:   "127.0.0.1",
			Port:     8187,
			LogLevel: "debug",
		},
	}

	data, err := json.MarshalIndent(sample, "", "  ")

	if err != nil {
		return fmt.Errorf("generate default config: %w", err)
	}

	os.MkdirAll("config", 0755)
	err = os.WriteFile(filepath.Join("config", "config.json"), data, 0644)
	if err != nil {
		return fmt.Errorf("generate default config: %w", err)
	}

	return nil
}
