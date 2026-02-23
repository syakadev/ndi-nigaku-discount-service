package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
)

type AppLogger struct {
	log zerolog.Logger
}

func NewAppLoggerDailyLog(logLevel zerolog.Level) *AppLogger {
	env := os.Getenv("SERVER_LOGS_ENV")
	isKubernetes := env == "Production"

	var writer io.Writer

	if isKubernetes {
		writer = os.Stdout
	} else {
		_ = os.MkdirAll("logs", 0755)

		today := time.Now().Format("02-01-2006") // DD-MM-YYYY
		logPath := filepath.Join("logs", fmt.Sprintf("%s.log", today))

		logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(fmt.Sprintf("Gagal membuka file log: %v", err))
		}

		consoleWriter := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}

		writer = io.MultiWriter(consoleWriter, logFile)
	}

	logger := zerolog.New(writer).
		With().
		Timestamp().
		Logger().
		Level(logLevel)

	return &AppLogger{log: logger}
}

// Logging helpers

func (l *AppLogger) Info(msg string) {
	l.log.Info().Msg(msg)
}

func (l *AppLogger) Error(msg string, err error) {
	l.log.Error().Err(err).Msg(msg)
}

func (l *AppLogger) Debug(msg string) {
	l.log.Debug().Msg(msg)
}

func (l *AppLogger) Warn(msg string, err error) {
	l.log.Warn().Err(err).Msg(msg)
}

func (l *AppLogger) Fatal(msg string) {
	l.log.Fatal().Msg(msg)
}

// WithFields membuat event log dengan field tambahan, bisa disambung .Msg(...)
func (l *AppLogger) WithFields(level zerolog.Level, fields map[string]interface{}) *zerolog.Event {
	event := l.log.WithLevel(level)
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	return event
}
