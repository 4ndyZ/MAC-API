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

func (l *Logger) Initialize(f string) {
	// Create console writer
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	// Create file writer
	fileWriter := &lumberjack.Logger{
		Filename:   f,
		MaxSize:    100, // Megabytes
		MaxBackups: 3,
		MaxAge:     28, // Days: true,
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

func (l *Logger) EnableDebug(d bool) {
	if l.Logger != nil {
		if d {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
			l.Logger.Info().Str("loglevel", "debug").Msg("LogLevel changed.")
		} else {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
			l.Logger.Info().Str("loglevel", "info").Msg("LogLevel changed.")
		}
	}
}

func (l *Logger) Rotate() {
	l.fileWriter.Rotate()
}
