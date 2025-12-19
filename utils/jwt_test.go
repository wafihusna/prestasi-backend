package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	token, err := GenerateJWT(
		"user-id",
		"student",
		[]string{}, // permissions (WAJIB)
	)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}