package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Тест на чтение конфигуратором несуществующего файла и значений по-умолчанию.
func TestConfigurator(t *testing.T) {
	config, err := readConfig("invalid_file_path")
	assert.Empty(t, config, "*Configuration must be nil")
	assert.Error(t, err, "open invalid_file_path: no such file or directory")
}
