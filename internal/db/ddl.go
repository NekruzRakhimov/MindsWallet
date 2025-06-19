package db

import "MindsWallet/logger"

func InitMigrations() error {
	usersTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			full_name VARCHAR(255) NOT NULL,
			username VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP,
			deleted_at TIMESTAMP
		);`

	_, err := db.Exec(usersTableQuery)
	if err != nil {
		logger.Error.Printf("[db] InitMigrations(): error during create users table: %v", err.Error())
		return err
	}

	accountsTableQuery := `
		CREATE TABLE IF NOT EXISTS accounts (
			id SERIAL PRIMARY KEY,
			user_id INT REFERENCES users(id),
			balance FLOAT DEFAULT 0,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP,
			deleted_at TIMESTAMP
		);`

	_, err = db.Exec(accountsTableQuery)
	if err != nil {
		logger.Error.Printf("[db] InitMigrations(): error during create accounts table: %v", err.Error())
		return err
	}
	return nil
}
