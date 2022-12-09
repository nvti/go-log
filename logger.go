package log

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type Log struct {
	log       zerolog.Logger
	skipFrame int
}

func New(level ...Level) *Log {
	log := &Log{
		skipFrame: 1,
	}

	var l Level
	if len(level) > 0 {
		l = level[0]
	} else {
		envLevel := strings.ToLower(os.Getenv("LOG_LEVEL"))
		l = Level(envLevel)
	}
	log.log = zerolog.New(formattedOutput(os.Stdout, true)).With().Timestamp().Logger().Level(convertLogLevel(l))

	return log
}

func (l *Log) Level(level Level) *Log {
	lvl := convertLogLevel(level)
	l.log = l.log.Level(lvl)

	return l
}

func (l *Log) SkipFrame(skipFrame int) *Log {
	l.skipFrame = skipFrame + 1

	return l
}

func (l *Log) LogFile(file string) *Log {
	var writers = []io.Writer{formattedOutput(os.Stdout, true)}

	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		_ = fmt.Errorf("Can't write log to file %s", file)
		return l
	}

	writers = append(writers, formattedOutput(f, false))
	mw := io.MultiWriter(writers...)

	l.log = l.log.Output(mw)
	return l
}

func (l *Log) Debug(v ...interface{}) {
	l.log.Debug().Msg(fmt.Sprint(v...))
}

func (l *Log) Info(v ...interface{}) {
	l.log.Info().Msg(fmt.Sprint(v...))
}

func (l *Log) Warn(v ...interface{}) {
	l.log.Warn().Msg(fmt.Sprint(v...))
}

func (l *Log) Error(v ...interface{}) {
	l.log.Error().CallerSkipFrame(l.skipFrame).Caller().Msg(fmt.Sprint(v...))
}

func (l *Log) Err(err error) {
	l.log.Err(err).CallerSkipFrame(l.skipFrame).Caller().Msg("")
}

func (l *Log) Fatal(v ...interface{}) {
	l.log.Fatal().CallerSkipFrame(l.skipFrame).Caller().Msg(fmt.Sprint(v...))
}

func (l *Log) Panic(v ...interface{}) {
	l.log.Panic().CallerSkipFrame(l.skipFrame).Caller().Msg(fmt.Sprint(v...))
}

func formattedOutput(output io.Writer, color bool) io.Writer {
	if isProduction {
		return output
	} else {
		return zerolog.ConsoleWriter{Out: output, TimeFormat: time.RFC3339, NoColor: !color}
	}
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
