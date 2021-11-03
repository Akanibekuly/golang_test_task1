package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Akanibekuly/golang_test_task1.git/internal/config"
	"github.com/Akanibekuly/golang_test_task1.git/internal/server"
	"github.com/joho/godotenv"
)

const (
	localEnvFilePath = "cmd/.env.local"
	prodEnvFilePath  = "cmd/.env"
)

func init() {
	if err := godotenv.Load(localEnvFilePath, prodEnvFilePath); err != nil {
		log.Fatalf("no .env file found: %s\n", err.Error())
	}
}

func Start() {
	conf := config.GetConfig()

	app := new(server.App)
	app.Initialize(conf)

	ctx, shutdown := context.WithCancel(context.Background())

	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		shutdown()
	}()

	app.Run(ctx)

}
