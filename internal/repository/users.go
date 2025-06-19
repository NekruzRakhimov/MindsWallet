package repository

import (
	"MindsWallet/internal/db"
	"MindsWallet/internal/models"
	"MindsWallet/logger"
)

func GetUserByUsernameAndPassword(username string, password string) (user models.User, err error) {
	err = db.GetDBConn().Get(&user, `SELECT id, 
					   full_name, 
					   username, 
					   created_at
				FROM users 
				WHERE deleted_at IS NULL 
				  AND username = $1
				  AND password = $2`, username, password)
	if err != nil {
		logger.Error.
			Printf("[repository] GetUserByUsernameAndPassword(): error duriing getting from database: %s\n", err.Error())
		return models.User{}, translateError(err)
	}

	return user, nil
}

func GetUserByUsername(username string) (user models.User, err error) {
	err = db.GetDBConn().Get(&user, `SELECT id, 
					   full_name, 
					   username, 
					   created_at
				FROM users WHERE deleted_at IS NULL AND username = $1`, username)
	if err != nil {
		logger.Error.
			Printf("[repository] GetUserByUsername(): error duriing getting from database: %s\n", err.Error())
		return models.User{}, translateError(err)
	}

	return user, nil
}

func CreateUser(user models.User) error {
	_, err := db.GetDBConn().Exec(`
			INSERT INTO users (full_name, username, password)
			VALUES ($1, $2, $3)`, user.FullName, user.Username, user.Password)
	if err != nil {
		logger.Error.
			Printf("[repository] CreateUser(): error duriing creating user from database: %s\n", err.Error())
		return translateError(err)
	}

	return nil
}
