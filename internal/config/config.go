package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
)

const (
	filename = "rssgopherconfig.json"
)

type Config struct {
	DatabaseURL string `json:"db_url"`
	Username    string `json:"current_user_name"`
}

func getJsonPath() (Path string) {
	currentDir, err := os.Getwd()
	if err != nil {
		return ""
	}

	return path.Join(currentDir, filename)
}

func openConfigFile(path string) (*os.File, error) {
	content, err := os.Open(path)
	if err != nil {
		return nil, errors.New("error opening config file: " + path)
	}

	return content, nil
}

func write(configInstance *Config, path string) error {
	data, err := json.MarshalIndent(configInstance, "", " ")
	if err != nil {
		return fmt.Errorf("error trying to convert Config struct to bytes: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("error writing config to file: %w", err)
	}

	return nil
}

func Read() (*Config, error) {
	jsonpath := getJsonPath()
	file, err := openConfigFile(jsonpath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, errors.New("error decoding config file")
	}

	return &config, nil
}

func (c *Config) SetUser(username string) error {
	c.Username = username
	jsonpath := getJsonPath()
	return write(c, jsonpath)
}
