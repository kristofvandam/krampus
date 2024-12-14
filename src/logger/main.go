package logger

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {
	lvl := new(slog.LevelVar)
	lvl.Set(slog.LevelInfo)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: lvl,
	}))

	if os.Getenv("DEBUG") == "true" {
		lvl.Set(slog.LevelDebug)
	}

	return logger
}
