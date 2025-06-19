package repository

import (
	"MindsWallet/internal/db"
	"MindsWallet/internal/models"
)

func GetAllAccounts() ([]models.Account, error) {
	var accounts []models.Account
	err := db.GetDBConn().Select(&accounts,
		`SELECT id, 
					   user_id, 
					   balance, 
					   created_at
				FROM accounts WHERE deleted_at IS NULL`)
	return accounts, err
}

func GetAccountByID(accountID int) (models.Account, error) {
	a := models.Account{}
	err := db.GetDBConn().Get(&a, `SELECT id, 
					   user_id, 
					   balance, 
					   created_at
				FROM accounts WHERE deleted_at IS NULL AND id = $1`, accountID)
	if err != nil {
		return models.Account{}, translateError(err)
	}

	return a, nil
}

func GetAccountByIDAndUserID(accountID, userID int) (models.Account, error) {
	a := models.Account{}
	err := db.GetDBConn().Get(&a, `SELECT id, 
					   user_id, 
					   balance, 
					   created_at
				FROM accounts 
				WHERE deleted_at IS NULL 
				  AND id = $1 AND user_id = $2`, accountID, userID)
	if err != nil {
		return models.Account{}, translateError(err)
	}

	return a, nil
}

func TopUpAccount(accountID int, amount float64) error {
	_, err := db.GetDBConn().Exec(`
			UPDATE accounts 
			SET balance = balance+$1, updated_at = CURRENT_TIMESTAMP
			WHERE id = $2`, amount, accountID)
	return err
}

func WithdrawFromAccount(accountID int, amount float64) error {
	_, err := db.GetDBConn().Exec(`
			UPDATE accounts 
			SET balance = balance-$1, updated_at = CURRENT_TIMESTAMP
			WHERE id = $2`, amount, accountID)
	return err

}
