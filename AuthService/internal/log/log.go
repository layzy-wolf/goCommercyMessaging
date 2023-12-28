package log

import (
	"github.com/mhchlib/go-kit/log"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func SetupLogger(env string) log.Logger {
	var logger log.Logger

	switch env {
	case envLocal:
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	case envDev:
		logger = log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	case envProd:
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	}

	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	return logger
}
