package service

import (
	"MindsWallet/internal/errs"
	"MindsWallet/internal/models"
	"MindsWallet/internal/repository"
	"MindsWallet/utils"
	"errors"
)

func CreateUser(u models.User) error {
	// 1. Проверить существует ли пользователь с таким username
	_, err := repository.GetUserByUsername(u.Username)
	if err != nil && !errors.Is(err, errs.ErrNotFound) {
		return err
	} else if err == nil {
		return errs.ErrUserAlreadyExists
	}

	// 2. Захешировать пароль
	u.Password = utils.GenerateHash(u.Password)

	// 3. Создаем пользователя
	if err = repository.CreateUser(u); err != nil {
		return err
	}

	return nil
}

func GetUserByUsernameAndPassword(username string, password string) (models.User, error) {
	// 1. Хешируем пароль
	password = utils.GenerateHash(password)

	// 2. Отправляем запрос в бд
	user, err := repository.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return models.User{}, errs.ErrIncorrectUsernameOrPassword
		}
		return models.User{}, err
	}
	return user, nil
}
