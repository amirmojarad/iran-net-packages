package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConfToml(t *testing.T) {
	config := GetConfToml()
	assert.NotEmpty(t, config.Providers.Irancell)
	assert.NotNil(t, config.Databases.Postgres)
	assert.NotEmpty(t, config.Databases.Postgres.Name)
	assert.NotEmpty(t, config.Databases.Postgres.User)
	assert.NotEmpty(t, config.Databases.Postgres.TimeZone)
	assert.NotEmpty(t, config.Databases.Postgres.Password)
	assert.NotEmpty(t, config.Databases.Postgres.Host)
	assert.NotZero(t, config.Databases.Postgres.Port)
	t.Log(config.Databases.Postgres)
}
