package main

import (
	"log/slog"

	logger "github.com/FelipePn10/dumb/config"
)

func main() {
	logger.InitLogger()
	slog.Info("Application started")
}
