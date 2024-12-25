package main

import (
	"fmt"

	"github.com/xclamation/go-sso-service/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: инициализировать логгер

	// TODO: инициализировать приложение (app, лежит отдельно)

	// TODO: запустить gRPC-сервер приложения

}
