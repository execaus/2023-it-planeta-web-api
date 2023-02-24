package main

import (
	"2023-it-planeta-web-api/configs"
	"2023-it-planeta-web-api/handler"
	"2023-it-planeta-web-api/models"
	"2023-it-planeta-web-api/repository"
	"2023-it-planeta-web-api/server"
	"2023-it-planeta-web-api/service"
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	var serverInstance server.Server

	env := models.LoadEnv()
	config := configs.LoadConfig()

	database := repository.NewBusinessDatabase(env, config)

	repos := repository.NewRepository(database)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	go runServer(&serverInstance, handlers, config.Server)

	runChannelStopServer()

	serverInstance.Shutdown(context.Background())
}

func runServer(server *server.Server, handlers *handler.Handler, config *configs.ServerConfig) {
	ginEngine := handlers.InitRoutes()

	if err := server.Run(config.Port, ginEngine); err != nil {
		if err.Error() != "http: Server closed" {
			logrus.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}
}

func runChannelStopServer() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGABRT)
	<-quit
}
