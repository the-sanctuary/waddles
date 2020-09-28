package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetToken(t *testing.T) {
	token := getToken("test.token")

	assert.Equal(t, "blahblahtesting123", token)
}

func TestGetTokenSad(t *testing.T) {
	token := getToken(".gitignore")

	assert.Equal(t, "", token)
}
