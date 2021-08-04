package zlog

import (
	"io"
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const service = "demo"

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:        os.Stderr,
			NoColor:    false,
			TimeFormat: "2006-01-02 15:04:05.999",
		},
	)
}

func InitConfig(logdev, loglev int, logfile string, maxsize, maxrotate, maxdays int) {
	SetLevel(loglev)
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05.999"
	zerolog.CallerSkipFrameCount = 4
	if logdev == 1 {
		log.Logger = log.Output(
			io.MultiWriter(zerolog.ConsoleWriter{
				Out:        os.Stderr,
				NoColor:    false,
				TimeFormat: "2006-01-02 15:04:05.999",
			}, &lumberjack.Logger{
				Filename:   logfile,
				MaxSize:    maxsize, // megabytes
				MaxBackups: maxrotate,
				MaxAge:     maxdays, // days
				Compress:   true,    // disabled by default
			}),
		).With().Caller().Logger()
	} else {
		log.Logger = log.Output(
			io.MultiWriter(zerolog.ConsoleWriter{
				Out:        os.Stderr,
				NoColor:    false,
				TimeFormat: "2006-01-02 15:04:05.999",
			}, &lumberjack.Logger{
				Filename:   logfile,
				MaxSize:    maxsize, // megabytes
				MaxBackups: maxrotate,
				MaxAge:     maxdays, // days
				Compress:   true,    // disabled by default                // disabled by default
			}),
		).With().Logger()
	}

}

func SetLevel(level int) {
	zerolog.SetGlobalLevel(zerolog.Level(level))
}

func Debugf(s string, v ...interface{}) {
	logEvent(log.Debug(), &s, &v)
}
func Infof(s string, v ...interface{}) {
	logEvent(log.Info(), &s, &v)
}
func Warnf(s string, v ...interface{}) {
	logEvent(log.Warn(), &s, &v)
}
func Errorf(s string, v ...interface{}) {
	logEvent(log.Error(), &s, &v)
}
func Fatalf(s string, v ...interface{}) {
	logEvent(log.Fatal(), &s, &v)
}

func logEvent(e *zerolog.Event, s *string, v *[]interface{}) {
	if v != nil {
		e.Msgf(*s, (*v)...)
	} else {
		e.Msgf(*s)
	}
}

func Debug() *zerolog.Event {
	return log.Debug().Str("service", service)
}

func Info() *zerolog.Event {
	return log.Info().Str("service", service)
}

func Warn() *zerolog.Event {
	return log.Warn().Str("service", service)
}

func Error() *zerolog.Event {
	return log.Error().Str("service", service)
}

func Fatal() *zerolog.Event {
	return log.Fatal().Str("service", service)
}
