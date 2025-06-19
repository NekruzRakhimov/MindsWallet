package controller

import (
	"MindsWallet/internal/errs"
	"MindsWallet/internal/models"
	"MindsWallet/internal/service"
	"MindsWallet/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Server is up and running",
	})
}

func GetAllAccounts(c *gin.Context) {
	accounts, err := service.GetAllAccounts()
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, accounts)
}

func GetAccountByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errs.ErrValidationFailed)
		HandleError(c, err)
		return
	}

	account, err := service.GetAccountByID(id)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, account)
}

func UpdateAccountBalance(c *gin.Context) {
	userID := c.GetInt(userIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUserIDNotFound)
		logger.Error.Printf("[controller] UpdateAccountBalance(): error during getting userID from conext: %s\n", errs.ErrUserIDNotFound)
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errs.ErrValidationFailed)
		HandleError(c, err)
		return
	}

	var balanceOperation models.BalanceOperation
	if err = c.ShouldBindJSON(&balanceOperation); err != nil {
		err = errors.Join(err, errors.New("invalid request body"), errs.ErrValidationFailed)
		HandleError(c, err)
		return
	}

	switch balanceOperation.Type {
	case models.TopUp:
		err = service.TopUpAccount(id, balanceOperation.Amount)
		if err != nil {
			HandleError(c, err)
			return
		}
	case models.Withdraw:
		err = service.WithdrawFromAccount(userID, id, balanceOperation.Amount)
		if err != nil {
			HandleError(c, err)
			return
		}
	default:
		HandleError(c, errs.ErrInvalidOperationType)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "account balance updated successfully",
	})

}
