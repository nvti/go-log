package log

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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

var (
	_log         *Log
	isProduction bool = false
)

func init() {
	production := os.Getenv("PRODUCTION")
	if production != "" && strings.ToLower(production) != "false" {
		isProduction = true
	} else {
		isProduction = false
		fmt.Println("Warning: Log is using development environment. To using production environment, set env PRODUCTION=true or call log.UseProduction()")
	}

	// print only relative path
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		short := file

		if cwd, err := os.Getwd(); err == nil {
			if rel, err := filepath.Rel(cwd, file); err == nil {
				short = rel
			}
		}
		file = short
		return file + ":" + strconv.Itoa(line)
	}

	_log = New().SkipFrame(1)
}

func UseProduction() {
	isProduction = true
	_log = New().SkipFrame(1)
}

func SetLevel(l Level) {
	_log = New(Level(l)).SkipFrame(1)
}

func LogFile(file string) {
	_log = _log.LogFile(file)
}

func Debug(v ...interface{}) {
	_log.Debug(v...)
}

func Info(v ...interface{}) {
	_log.Info(v...)
}

func Warn(v ...interface{}) {
	_log.Warn(v...)
}

func Error(msg string) {
	_log.Error(msg)
}

func Err(err error) {
	_log.Err(err)
}

func Fatal(msg string) {
	_log.Fatal(msg)
}

func Panic(msg string) {
	_log.Panic(msg)
}
