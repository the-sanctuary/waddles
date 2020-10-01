package util

import (
	"bufio"
	"os"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	ConfigFile string = "./waddles.toml"
)

var (
	//Cfg holds the current config information in a Config struct
	Cfg Config
)

// ConfigDatabase holds bot config information
type Config struct {
	Wadl struct {
		LogLevel string `toml:"log-level" env:"WADL_DEBUG" env-default:"info"`
		Prefix   string `toml:"prefx" env:"" env-default:"!"`
		Token    string `toml:"token" env:"WADL_TOKEN" env-default:""`
	} `toml:"wadl"`
	Db struct {
		User string `toml:"user" env-default:"waddles"`
		Pass string `toml:"pass" env-default:""`
		Host string `toml:"host" env-default:"localhost"`
		Port string `toml:"port" env-default:"5432"`
		URL  string `toml:"url" env:"DATABASE_URL"`
		Name string `toml:"name" env-default:"waddles"`
	} `toml:"db"`
}

//ReadConfig parses config options from the environment and config file into a ConfigDatabase struct
func ReadConfig() {
	// Read in environment variables, and set the log level
	err := cleanenv.ReadConfig(ConfigFile, &Cfg)

	if err != nil {
		log.Info().Msgf("[CONF] Unable to read  config file: \"%s\".  Continuing with defaults.", ConfigFile)
	}

	err = cleanenv.ReadEnv(&Cfg)

	if err != nil {
		log.Info().Msg("[CONF] Unable to read in environment variables.")
	}

	//parse zerolog.Level from Cfg.Debug
	globalLevel, err := zerolog.ParseLevel(Cfg.Wadl.LogLevel)

	log.Info().Msgf("[LOG] Log Level set to: %s", globalLevel.String())

	if err != nil {
		log.Info().Msgf("[CONF] Supplied debugging log level (%d) is invalid. Defaulting to \"info\".", Cfg.Wadl.LogLevel)
	}

	// Set global log level
	zerolog.SetGlobalLevel(globalLevel)
}

func parseLogLevel(lglvl int) zerolog.Level {
	var final zerolog.Level = zerolog.InfoLevel
	switch lglvl {
	case 0:
		final = zerolog.WarnLevel
		break
	case 1:
		final = zerolog.InfoLevel
		break
	case 2:
		final = zerolog.DebugLevel
		break
	default:
		final = zerolog.InfoLevel
		break
	}
	return final
}

func getToken(filepath string) (string, error) {
	file, err := os.Open(filepath)

	if err != nil {
		return "", err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.Split(scanner.Text(), "=")[0] == "token" {
			return strings.Split(scanner.Text(), "=")[1], nil
		}
	}
	return "", err
}
