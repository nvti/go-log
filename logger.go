package log

import (
	"os"

	"github.com/rs/zerolog"
)

func New(l Level) *Log {
	log = &Log{}
	level := convertLogLevel(l)
	log.log = zerolog.New(os.Stdout).With().Timestamp().Logger().Level(level)
	log.detailLog = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger().Level(level)

	return log
}

func (l *Log) Debug(msg string) {
	l.log.Debug().Msg(msg)
}

func (l *Log) Error(msg string) {
	l.detailLog.Error().Msg(msg)
}

func convertLogLevel(level Level) zerolog.Level {
	switch level {
	case DebugLevel:
		return zerolog.DebugLevel
	case InfoLevel:
		return zerolog.InfoLevel
	case WarnLevel:
		return zerolog.WarnLevel
	case ErrorLevel:
		return zerolog.ErrorLevel
	case FatalLevel:
		return zerolog.FatalLevel
	case PanicLevel:
		return zerolog.PanicLevel
	default:
		return zerolog.InfoLevel
	}
}
