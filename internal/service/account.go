package service

import (
	"MindsWallet/internal/errs"
	"MindsWallet/internal/models"
	"MindsWallet/internal/repository"
	"errors"
)

func GetAllAccounts() ([]models.Account, error) {
	accounts, err := repository.GetAllAccounts()
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func GetAccountByID(accountID int) (models.Account, error) {
	account, err := repository.GetAccountByID(accountID)
	if err != nil {
		return models.Account{}, err
	}

	return account, nil
}

func TopUpAccount(accountID int, amount float64) error {
	// проверить существует ли такой кошелек
	_, err := repository.GetAccountByID(accountID)
	if err != nil {
		return err
	}

	// отправить запрос на пополнение
	err = repository.TopUpAccount(accountID, amount)
	if err != nil {
		return err
	}

	return nil
}

func WithdrawFromAccount(userID, accountID int, amount float64) error {
	// проверить существует ли такой кошелек
	account, err := repository.GetAccountByID(accountID)
	if err != nil {
		return err
	}

	// проверить принадлежит ли этот кошелек указанному пользователю
	_, err = repository.GetAccountByIDAndUserID(accountID, userID)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return errs.ErrNoPermissionsToWithdraw
		}
	}

	// проверить хватит ли денег
	if account.Balance < amount {
		return errs.ErrNotEnoughBalance
	}

	// отправить запрос на снятие
	err = repository.WithdrawFromAccount(accountID, amount)
	if err != nil {
		return err
	}

	// ответить клиенту
	return nil
}
