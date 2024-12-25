package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/xclamation/go-sso-service/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// DID: инициализировать конфиг
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// DID: инициализировать логгер
	log := setupLogger(cfg.Env)

	log.Info("starting application", slog.Any("cfg", cfg)) // It is better to NOT log config due to it's sensitive data
	
	// TODO: реализовать красивый логгер с цветной подсветкой

	// TODO: инициализировать приложение (app, лежит отдельно)

	// TODO: запустить gRPC-сервер приложения

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
