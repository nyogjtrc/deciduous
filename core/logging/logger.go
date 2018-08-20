package logging

import "go.uber.org/zap"

var logger *zap.Logger

func init() {
	newConsoleFileLogger()
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

// L return zap logger
func L() *zap.Logger {
	return logger
}
