package logging

import "go.uber.org/zap"

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
