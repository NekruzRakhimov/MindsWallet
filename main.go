package main

import (
	"MindsWallet/internal/configs"
	"MindsWallet/internal/controller"
	"MindsWallet/internal/db"
	"MindsWallet/internal/server"
	"MindsWallet/logger"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Reading configs
	if err := configs.ReadSettings(); err != nil {
		log.Fatalf("Ошибка чтения настроек: %s", err)
	}

	// Initializing logger
	if err := logger.Init(); err != nil {
		return
	}
	logger.Info.Println("Loggers initialized successfully!")

	// Connecting to db
	if err := db.ConnectDB(); err != nil {
		return
	}
	logger.Info.Println("Connection to database established successfully!")

	// Initializing db-migrations
	if err := db.InitMigrations(); err != nil {
		return
	}
	logger.Info.Println("Migrations initialized successfully!")

	// Running http-server
	//if err := controller.RunServer(); err != nil {
	//	return
	//}

	mainServer := new(server.Server)
	go func() {
		if err := mainServer.Run(configs.AppSettings.AppParams.PortRun, controller.RunServer()); err != nil {
			log.Fatalf("Ошибка при запуске HTTP сервера: %s", err)
		}
	}()

	// Ожидание сигнала для завершения работы
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Printf("\nНачало завершения программ\n")

	// Закрытие соединения с БД, если необходимо
	if err := db.GetDBConn().Close(); err != nil {
		log.Fatalf("Ошибка при закрытии соединения с БД: %s", err)

	}
	fmt.Println("Соединение с БД успешно закрыто")

	// Используем контекст с тайм-аутом для завершения работы сервера
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := mainServer.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка при завершении работы сервера: %s", err)
	}

	fmt.Println("HTTP-сервис успешно выключен")
	fmt.Println("Конец завершения программы")

}
