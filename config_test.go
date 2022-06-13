package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	c := GetConfig()
	assert.Equal(t, "a", c.Database.dsn, "The two words should be the same.")
}
