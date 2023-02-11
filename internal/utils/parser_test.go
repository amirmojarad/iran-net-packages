package utils

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestParseAmount(t *testing.T) {
	raw := "                                                        15 گیگابایت"
	res, err := ParseAmount(raw)
	assert.Nil(t, err)
	assert.Equal(t, int(math.Pow(10, 9))*15, res)
}

func TestParseDuration(t *testing.T) {
	raw := "                                                         120 روزه"
	res, err := ParseDuration(raw)
	assert.Nil(t, err)
	assert.Equal(t, 120, res)
}

func TestParsePrice(t *testing.T) {
	raw := "                                                    62,000 تومان"
	res, err := ParsePrice(raw)
	assert.Nil(t, err)
	assert.Equal(t, 62000, res)
}
