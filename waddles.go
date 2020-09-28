package main

import (
	"bufio"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ConfigDatabase struct {
	Debug int `env:"WADDLESDEBUG" env-default:"1"`
}

var cfg ConfigDatabase
var token string

func getToken(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Print("[CONF] Unable to open token file for reading.  Quitting...")
		//log.Printf("[2ERR] %s.", err.Error())
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
		log.Print("[CONF] Unable to read in environment variables.  Continuing with defaults.")
	}
	log.Printf("[CONF] Debug Level: %d.", cfg.Debug)
	//zerolog.SetGlobalLevel(cfg.Debug)

	// Get filepath for token
	filepath := "./waddles.token"
	if len(os.Args) > 1 {
		filepath = os.Args[1]
	}
	getToken(filepath)

	// Create a Discord session using our bot token (client secret)
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Print("[WADL] Unable to create a Discord session.  Quitting....")
		os.Exit(1)
	}

	// Open a websocket connection to Discord and start listening
	err = dg.Open()
	if err != nil {
		log.Print("[WADL] Unable to open a connection to Discord.  Quitting....")
		log.Print(err)
		os.Exit(1)
	}
	defer dg.Close()

	// Print msg that the bot is running
	log.Print("[WADL] Waddles is now running.  Press CTRL-C to quit.")

	// Register term signals
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
