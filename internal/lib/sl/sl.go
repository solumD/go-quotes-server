package sl

import (
	"log/slog"
	"os"
)

const (
	levelLocal = "local"
	levelDev   = "dev"
	levelProd  = "prod"
)

// InitLogger initializes the slog logger.
func InitLogger(loggerLevel string) *slog.Logger {
	var log *slog.Logger

	switch loggerLevel {
	case levelLocal:
		log = slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug},
			),
		)
	case levelDev:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug},
			),
		)
	case levelProd:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelInfo},
			),
		)
	}

	return log
}

// Err returns an error attribute.
func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
