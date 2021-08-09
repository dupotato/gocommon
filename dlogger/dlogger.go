package dlogger

import (
	"gocommon/dlogger/grpclog"
	"gocommon/dlogger/zlog"
)

type Logger interface {
	Debugf(s string, v ...interface{})
	Infof(s string, v ...interface{})
	Warnf(s string, v ...interface{})
	Errorf(s string, v ...interface{})
	Fatalf(s string, v ...interface{})
}

var gLogger Logger
var gLogConfig LogConfig
var mapLogger map[string]*Logger

type LogConfig struct {
	loglev    int
	logdev    int
	filename  string
	maxsize   int
	maxrotate int
	maxages   int
}

func Register(name string, logg *Logger) {
	mapLogger[name] = logg
}

func init() {
	gLogConfig = LogConfig{
		loglev:    0,
		logdev:    0,
		filename:  "common.log",
		maxsize:   10,
		maxrotate: 10,
		maxages:   1,
	}
}

func Init(name string) {
	if name == "zero" {
		gLogger = zlog.NewZlog(gLogConfig.loglev, gLogConfig.logdev, gLogConfig.filename, gLogConfig.maxsize, gLogConfig.maxrotate, gLogConfig.maxages)
	} else if name == "grpc" {
		gLogger = grpclog.NewZapLogger()

	} else {
		gLogger = zlog.NewZlog(gLogConfig.loglev, gLogConfig.logdev,
			gLogConfig.filename, gLogConfig.maxsize, gLogConfig.maxrotate, gLogConfig.maxages)
	}
}

func Debugf(s string, v ...interface{}) {
	gLogger.Debugf(s, v)
}
func Infof(s string, v ...interface{}) {
	gLogger.Infof(s, v)
}
func Warnf(s string, v ...interface{}) {
	gLogger.Warnf(s, v)
}
func Errorf(s string, v ...interface{}) {
	gLogger.Errorf(s, v)
}
func Fatalf(s string, v ...interface{}) {
	gLogger.Fatalf(s, v)
}
