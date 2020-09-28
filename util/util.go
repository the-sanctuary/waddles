package util

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/snwfdhmp/errlog"
)

var (
	errorLogger errlog.Logger
)

//SetupLogging sets up errlog and zerolog and sets errlog to use zerolog to
func SetupLogging() {
	errorLogger = errlog.NewLogger(&errlog.Config{
		PrintFunc:          log.Error().Msgf,
		LinesBefore:        6,
		LinesAfter:         4,
		PrintError:         true,
		PrintSource:        true,
		PrintStack:         false,
		ExitOnDebugSuccess: true,
	})

	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123Z})
}

// DebugError handles an error with errlog (& zerolog)
func DebugError(err error) bool {
	return errorLogger.Debug(err)
}
