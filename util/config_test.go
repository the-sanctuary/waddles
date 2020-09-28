package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetToken(t *testing.T) {
	SetupLogging()
	token, err := getToken("../example.token")

	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, "<BOT-TOKEN-GOES-HERE>", token)
}

func TestGetTokenSad(t *testing.T) {
	SetupLogging()
	token, err := getToken("../gitignore")

	assert.NotNil(t, err, "err should not nil")
	assert.Equal(t, "", token)
}
