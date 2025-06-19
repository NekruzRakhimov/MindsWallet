package main

import (
	"MindsWallet/internal/configs"
	"MindsWallet/internal/controller"
	"MindsWallet/internal/db"
	"MindsWallet/logger"
	"log"
)

func main() {
	// Чтение настроек
	if err := configs.ReadSettings(); err != nil {
		log.Fatalf("Ошибка чтения настроек: %s", err)
	}

	// Инициализация логгера
	if err := logger.Init(); err != nil {
		return
	}
	logger.Info.Println("Loggers initialized successfully!")

	if err := db.ConnectDB(); err != nil {
		return
	}

	logger.Info.Println("Connection to database established successfully!")

	if err := db.InitMigrations(); err != nil {
		return
	}
	logger.Info.Println("Migrations initialized successfully!")

	if err := controller.RunServer(); err != nil {
		return
	}

}
