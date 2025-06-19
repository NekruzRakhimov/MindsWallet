package configs

import (
	"MindsWallet/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var AppSettings models.Configs

func ReadSettings() error {
	fmt.Println("Loading .env file")

	// Просто грузим .env в переменные окружения
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not found, using system environment variables")
	}

	fmt.Println("Reading settings file: configs/configs.json")
	configFile, err := os.Open("internal/configs/configs.json")
	if err != nil {
		return errors.New(fmt.Sprintf("Couldn't open config file: %s", err.Error()))
	}
	defer func(configFile *os.File) {
		err = configFile.Close()
		if err != nil {
			log.Fatal("Couldn't close config file: ", err.Error())
		}
	}(configFile)

	if err = json.NewDecoder(configFile).Decode(&AppSettings); err != nil {
		return errors.New(fmt.Sprintf("Couldn't decode json config: %s", err.Error()))
	}

	return nil
}
