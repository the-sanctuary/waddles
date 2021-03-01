package util

import (
	"math"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

//RegisterTermSignals  -
func RegisterTermSignals() {
	// Register term signals
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

//SliceContains returns true if str is included in slice
func SliceContains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

//AbsInt returns the absolute value of an integer
func AbsInt(i int) int {
	return int(math.Abs(float64(i)))
}

//FileExists returns true if the given path exists and isn't a directory
func FileExists(filename string) bool {
	s, err := os.Stat(filename)

	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		log.Error().Err(err).Msgf("helpers.FileExists(%s) errored out:", err) // If the err wasn't expected, something really went wrong
	} else if s == nil { // If no s is returned there is a different issue.
		return false
	}
	return !s.IsDir()
}

// ComesFromDM returns true if a message comes from a DM channel
func ComesFromDM(s *discordgo.Session, m *discordgo.MessageCreate) (bool, error) {
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		if channel, err = s.Channel(m.ChannelID); err != nil {
			return false, err
		}
	}

	return channel.Type == discordgo.ChannelTypeDM, nil
}
