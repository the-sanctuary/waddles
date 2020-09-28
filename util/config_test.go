package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetToken(t *testing.T) {
	token := util.getToken("example.token")

	assert.Equal(t, "<BOT-TOKEN-GOES-HERE>", token)
}

func TestGetTokenSad(t *testing.T) {
	token := util.getToken(".gitignore")

	assert.Equal(t, "", token)
}
