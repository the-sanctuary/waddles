package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	//Cfg holds the current config information in a ConfigDatabase struct
	Cfg ConfigDatabase
)

// ConfigDatabase holds bot config information
type ConfigDatabase struct {
	Debug     int    `env:"WADL_DEBUG" env-default:"1"`
	Token     string `env:"WADL_TOKEN" env-default:""`
	TokenFile string `env:"WADL_TOKEN" env-default:"./waddles.token"`
}

//ReadConfig parses config options from the environment and config file into a ConfigDatabase struct
func ReadConfig() {
	// Read in environment variables, and set the log level
	err := cleanenv.ReadEnv(&Cfg)

	if err != nil {
		log.Info().Msg("[CONF] Unable to read in environment variables.  Continuing with defaults.")
	}

	switch Cfg.Debug {
	case 0:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
		log.Info().Msg("[CONF] Debug Level: " + strconv.Itoa(Cfg.Debug) + ", Silent.")
		break
	case 1:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Info().Msg("[CONF] Debug Level: " + strconv.Itoa(Cfg.Debug) + ", Info.")
		break
	case 2:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Info().Msg("[CONF] Debug Level: " + strconv.Itoa(Cfg.Debug) + ", Debug.")
		break
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Info().Msg("[CONF] Debug Level: " + strconv.Itoa(Cfg.Debug) + ", Invalid.  Using Debug Level: 1, Info.")
		break
	}

	// If token wasn't provided via env/config file, read it from a token file
	if len(Cfg.Token) == 0 {
		// Get filepath for token
		filepath := "./waddles.token"

		if len(os.Args) > 1 {
			filepath = os.Args[1]
		}

		Cfg.Token, err = getToken(filepath)

		if err != nil {
			log.Info().Msg("[CONF] Unable to open token file for reading.  Quitting...")
			log.Debug().Msg("[IERR] " + err.Error())
			os.Exit(1)
		}
	}
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
