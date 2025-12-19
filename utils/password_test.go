package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashAndComparePassword(t *testing.T) {
	password := "secret123"

	hash, err := HashPassword(password)
	assert.NoError(t, err)

	ok := CheckPasswordHash(password, hash)
	assert.True(t, ok)
}