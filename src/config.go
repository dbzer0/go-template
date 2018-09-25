package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	LogFile *string `json:"logfile,omitempty"`
}

type ErrNoOptionFound string

func (op ErrNoOptionFound) Error() error {
	return fmt.Errorf("[ERROR] '%s' not found in configuration file", op)
}

// readConfig() - читает конфигурационный файл из JSON в структуру
// данных Configuration.
func readConfig(configFile string) (*Configuration, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var configuration *Configuration
	if err = json.NewDecoder(file).Decode(&configuration); err != nil {
		return nil, err
	}

	return configuration, err
}

// Проверяет указаны ли обязательные конфигуарционные опции.
func (c *Configuration) Validate() error {
	//...
	return nil
}
