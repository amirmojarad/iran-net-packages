package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConfToml(t *testing.T) {
	config := GetConfToml()
	assert.NotEmpty(t, config.Providers.Irancell)
}
