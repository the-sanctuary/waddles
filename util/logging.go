package util

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/snwfdhmp/errlog"
)

var errorLogger errlog.Logger

//InitializeLogging inits basic logging capabilities
func InitializeLogging() {
	//set pretty console output for zerolog
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC1123Z})
}

//SetupLogging sets up errlog and zerolog and sets errlog to use zerolog to
func SetupLogging() {
	errorLogger = errlog.NewLogger(&errlog.Config{
		PrintFunc:          log.Error().Msgf,
		LinesBefore:        6,
		LinesAfter:         4,
		PrintError:         true,
		PrintSource:        true,
		PrintStack:         true,
		ExitOnDebugSuccess: true,
	})
	if zerolog.GlobalLevel() <= zerolog.TraceLevel {
		//adds file and line number to log
		log.Logger = log.With().Caller().Logger()
	} else {
		errlog.DefaultLogger.Disable(true)
	}
}

// DebugError handles an error with errlog (& zerolog)
func DebugError(err error) bool {
	isNil := errorLogger.Debug(err)
	if !isNil {
		log.Error().Err(err)
	}
	return isNil
}
