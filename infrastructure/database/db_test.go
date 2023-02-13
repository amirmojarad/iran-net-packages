package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnect(t *testing.T) {
	handle := Connect()
	d, err := handle.DB()
	assert.Nil(t, err)
	err = d.Ping()
	assert.Nil(t, err)
}

func TestConnectToTest(t *testing.T) {
	handle := ConnectToTest()
	d, err := handle.DB()
	assert.Nil(t, err)
	err = d.Ping()
	assert.Nil(t, err)
}
