package cfg

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/pelletier/go-toml"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/the-sanctuary/waddles/pkg/util"
)

var cfg *Config

func Cfg() *Config {
	if cfg == nil {
		cfg = ReadConfig()
	}
	return cfg
}

//ReadConfig parses the config file into a Config struct
func ReadConfig() *Config {
	configDir := os.Getenv("WADL_CONFIG_DIR")

	if configDir == "" {
		pwd, _ := os.Getwd()
		configDir = pwd + "/config/"
		log.Warn().Msgf("WADL_CONFIG_DIR not set, defaulting to working dir (%s)", configDir)
	}

	if !strings.HasSuffix(configDir, "/") {
		configDir = path.Clean(configDir) + "/"
	}

	config := &Config{configDir: configDir}

	configFile := config.GetConfigFileLocation("waddles.toml")

	if !util.FileExists(configFile) {
		config.configDir = ""

		var bytes bytes.Buffer
		err := toml.NewEncoder(&bytes).Order(toml.OrderPreserve).Encode(config)

		if err != nil {
			log.Panic().Err(err).Msg("Unable to save sample config file.")
		}

		ioutil.WriteFile(configFile, bytes.Bytes(), 0644)
		log.Fatal().Msgf("Config file doesn't exist. An example has been saved in its place.")
	}

	// Read config file from the file
	bytes, err := ioutil.ReadFile(configFile)

	if err != nil {
		log.Fatal().Err(err).Msgf("Unable to read config file at: '%s'", configFile)
	}

	// Unmarshal the config file bytes into a Config struct
	err = toml.Unmarshal(bytes, config)

	if err != nil {
		log.Fatal().Err(err).Msg("Unable to parse config file.")
	}

	log.Debug().Msgf("Read config file: %s", configFile)
	log.Trace().Msgf("Config Struct: %+v", *config)

	logLevel, err := zerolog.ParseLevel(config.Wadl.LogLevel)
	if err != nil {
		log.Warn().Msgf("Supplied config file log level (%s) is invalid. Defaulting to info.", config.Wadl.LogLevel)
		logLevel = zerolog.InfoLevel
	}

	log.Info().Msgf("Log Level set to: %s", logLevel.String())

	// Set global log level
	zerolog.SetGlobalLevel(logLevel)

	return config
}
