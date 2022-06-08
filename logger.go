package main

import (
	"github.com/rs/zerolog"
	_ "github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

// Logger struct to hold refs
type Logger struct {
	Logger     *zerolog.Logger
	fileWriter *lumberjack.Logger
}

// Initialize logger with config file location as string
func (l *Logger) Initialize(f string) {
	// Create console writer
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	// Create file writer
	fileWriter := &lumberjack.Logger{
		Filename:   f,
		MaxSize:    10, // Megabytes
		MaxBackups: 5,
		MaxAge:     28, // Days
		LocalTime:  true,
	}
	multi := zerolog.MultiLevelWriter(consoleWriter, fileWriter)
	//
	logger := zerolog.New(multi).With().Timestamp().Logger()
	// Set default loglevel
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	// Set logger to struct
	l.fileWriter = fileWriter
	l.Logger = &logger
}

// EnableDebug enable the debugging mode of the logger depending on bool
func (l *Logger) EnableDebug(debug bool) {
	if l.Logger != nil {
		if debug {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
			l.Logger.Info().Str("loglevel", "debug").Msg("LogLevel changed.")
		} else {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
			l.Logger.Info().Str("loglevel", "info").Msg("LogLevel changed.")
		}
	}
}
