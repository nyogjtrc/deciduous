package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger
var slogger *zap.SugaredLogger

func init() {
	newConsoleFileLogger()
	newSugaredLogger()
}

func newDevelopment() {
	logger, _ = zap.NewDevelopment()
}

func newConsoleFileLogger() {
	var err error
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"stderr",
		"/tmp/service.log",
	}

	logger, err = cfg.Build()
	if err != nil {
		panic(err)
	}
}

func newSugaredLogger() {
	slogger = logger.Sugar()
}

// L return zap logger
func L() *zap.Logger {
	return logger
}

// S return zap sugared logger
func S() *zap.SugaredLogger {
	return slogger
}

// NewRollingLogger will rotate log file
func NewRollingLogger() *zap.Logger {
	writeFile := zapcore.AddSync(&lumberjack.Logger{
		Filename: "service.log",
		MaxSize:  10, // megabytes
		MaxAge:   28, // days
	})
	writeStdout := zapcore.AddSync(os.Stdout)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}),
		zapcore.NewMultiWriteSyncer(writeFile, writeStdout),
		zap.NewAtomicLevel(),
	)
	return zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	)
}
