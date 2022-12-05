package log

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)

// Level defines log levels.
type Level string

const (
	// DebugLevel defines debug log level.
	DebugLevel Level = "debug"
	// InfoLevel defines info log level.
	InfoLevel Level = "info"
	// WarnLevel defines warn log level.
	WarnLevel Level = "warn"
	// ErrorLevel defines error log level.
	ErrorLevel Level = "error"
	// FatalLevel defines fatal log level.
	FatalLevel Level = "fatal"
	// PanicLevel defines panic log level.
	PanicLevel Level = "panic"
)

type Log struct {
	log       zerolog.Logger
	detailLog zerolog.Logger
}

var (
	log *Log
)

func init() {
	production := os.Getenv("PRODUCTION")
	envLevel := strings.ToLower(os.Getenv("LOG_LEVEL"))
	if production != "" && strings.ToLower(production) != "false" {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	}

	log = New(Level(envLevel))
}

func Debug(msg string) {
	log.Debug(msg)
}

func Error(msg string) {
	log.Error(msg)
}
