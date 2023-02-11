package irancell

import (
	"github.com/stretchr/testify/assert"
	"internet_service/config"
	"testing"
)

func TestIrancell_Fetch(t *testing.T) {
	conf := config.GetConfToml()
	irancell := New(conf.Providers.Irancell)
	irancell.Fetch()
	assert.NotEmpty(t, irancell.Packages)
}
