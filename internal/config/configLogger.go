package config

import (
	"errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

type Log struct {
	logFile   string
	logToFile bool
	logLevel  zerolog.Level
}

func (l *Log) ConfigureLog() error {
	// logger
	ll, err := zerolog.ParseLevel(viper.GetString("log.level"))
	if err != nil {
		return errors.New("error parse log level")
	}
	l.logLevel = ll
	l.logFile = viper.GetString("log.filename")
	l.logToFile = viper.GetBool("log.to_file")
	l.configureLogger()
	return nil
}

func (l *Log) configureLogger() {
	zerolog.SetGlobalLevel(l.logLevel)
	if l.logToFile {
		z := zerolog.New(&lumberjack.Logger{
			Filename:   l.logFile,
			MaxSize:    10,
			MaxBackups: 2,
			MaxAge:     10,
			Compress:   true,
		}).With().Timestamp().Logger()
		log.Logger = z
	} else {
		z := zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			NoColor:    false,
			TimeFormat: time.UnixDate,
		}).With().Timestamp().Logger()
		log.Logger = z
	}
}

func (l *Log) LogLevel() zerolog.Level {
	return l.logLevel
}

func (l *Log) LogFileName() string {
	return l.logFile
}
func (l *Log) LogToFile() bool {
	return l.logToFile
}

func (l *Log) LogConfig() Log {
	return *l
}
