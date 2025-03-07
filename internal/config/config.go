package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const filePath = "/bootdev/postnest/postnestconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	file_path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	json_file, err := os.Open(file_path)
	if err != nil {
		return Config{}, fmt.Errorf("error reading file %s: %v", file_path, err)
	}
	defer json_file.Close()

	decoder := json.NewDecoder(json_file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username

	data, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("error converting config to json: %v", err)
	}

	file_path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	writeJSONToFile(file_path, data)

	return nil
}

func writeJSONToFile(filename string, jsonData []byte) error {
	// Open the file with write mode and truncate the file (replace contents)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error opening/creating file: %v", err)
	}
	defer file.Close()

	// Write the JSON data to the file
	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("error writing JSON to file: %v", err)
	}

	return nil
}

func getConfigFilePath() (string, error) {
	home_path, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home directory: %v", err)
	}

	file_path := filepath.Join(home_path, filePath)
	return file_path, nil
}
