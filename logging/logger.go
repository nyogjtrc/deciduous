package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const filename string = "service.log"

// NewConsoleFileLogger will output log to stderr and file
func NewConsoleFileLogger() *zap.Logger {
	var err error
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"stderr",
		filename,
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return logger
}

// NewRollingLogger will rotate log file
func NewRollingLogger() *zap.Logger {
	writeFile := zapcore.AddSync(&lumberjack.Logger{
		Filename: filename,
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
