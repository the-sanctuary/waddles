package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"github.com/the-sanctuary/waddles/handlers"
	"github.com/the-sanctuary/waddles/util"
)

func main() {
	util.InitializeLogging()
	util.ReadConfig()
	util.SetupLogging()

	// Create a Discord session using our bot token (client secret)
	dg, err := discordgo.New("Bot " + util.Cfg.Token)
	if util.DebugError(err) {
		log.Info().Msg("[WADL] Unable to create a Discord session.  Quitting....")
		log.Debug().Msg("[IERR] " + err.Error())
		os.Exit(1)
	}

	// Register handlers
	dg.AddHandler(handlers.MsgHandler)

	// Open a websocket connection to Discord and start listening
	err = dg.Open()
	if util.DebugError(err) {
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
