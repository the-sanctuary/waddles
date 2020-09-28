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
	//set pretty console output for zerolog
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123Z})

	if Cfg.Debug == 2 {
		errorLogger = errlog.NewLogger(&errlog.Config{
			PrintFunc:          log.Error().Msgf,
			LinesBefore:        6,
			LinesAfter:         4,
			PrintError:         true,
			PrintSource:        true,
			PrintStack:         false,
			ExitOnDebugSuccess: true,
		})

		//adds file and line number to log
		log.Logger = log.With().Caller().Logger()
	} else {
		errorLogger = errlog.DefaultLogger
		errlog.DefaultLogger.Disable(true)
	}
}

// DebugError handles an error with errlog (& zerolog)
func DebugError(err error) bool {
	return errorLogger.Debug(err)
}
