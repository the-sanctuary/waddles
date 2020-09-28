package main

import (
	"bufio"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ConfigDatabase struct {
	Debug int    `env:"WADLDEBUG" env-default:"1"`
	Token string `env:"WADLTOKEN" env-default:""`
}

var cfg ConfigDatabase
var token string

func getToken(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Info().Msg("[CONF] Unable to open token file for reading.  Quitting...")
		log.Debug().Msg("[IERR] " + err.Error())
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		token = strings.Split(scanner.Text(), "=")[1]
	}
}

func main() {
	// Set output for zerolog
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Read in environment variables, and set the log level
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Info().Msg("[CONF] Unable to read in environment variables.  Continuing with defaults.")
	}
	switch cfg.Debug {
	case 0:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
		log.Info().Msg("[CONF] Debug Level: " + strconv.Itoa(cfg.Debug) + ", Silent.")
		break
	case 1:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Info().Msg("[CONF] Debug Level: " + strconv.Itoa(cfg.Debug) + ", Info.")
		break
	case 2:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Info().Msg("[CONF] Debug Level: " + strconv.Itoa(cfg.Debug) + ", Debug.")
		break
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Info().Msg("[CONF] Debug Level: " + strconv.Itoa(cfg.Debug) + ", Invalid.  Using Debug Level: 1, Info.")
		break
	}

	// Get the token from an environment variable or a token file
	if len(cfg.Token) > 0 {
		token = cfg.Token
	} else {
		// Get filepath for token
		filepath := "./waddles.token"
		if len(os.Args) > 1 {
			filepath = os.Args[1]
		}
		getToken(filepath)
	}

	// Create a Discord session using our bot token (client secret)
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Info().Msg("[WADL] Unable to create a Discord session.  Quitting....")
		log.Debug().Msg("[IERR] " + err.Error())
		os.Exit(1)
	}

	// Open a websocket connection to Discord and start listening
	err = dg.Open()
	if err != nil {
		log.Info().Msg("[WADL] Unable to open a connection to Discord.  Quitting....")
		log.Debug().Msg("[IERR] " + err.Error())
		os.Exit(1)
	}
	defer dg.Close()

	// Print msg that the bot is running
	log.Info().Msg("[WADL] Waddles is now running.  Press CTRL-C to quit.")

	// Register term signals
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
