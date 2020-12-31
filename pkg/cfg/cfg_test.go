package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetConfigFileLocation(t *testing.T) {
	config1 := Config{
		configDir: "/foo/bar/",
	}
	config2 := Config{
		configDir: "/foo/bar/haz",
	}

	assert.Equal(t, "/foo/bar/config1.toml", config1.GetConfigFileLocation("config1.toml"))
	assert.Equal(t, "/foo/bar/haz/config2.toml", config2.GetConfigFileLocation("config2.toml"))
}
