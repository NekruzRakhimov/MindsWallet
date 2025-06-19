package controller

import (
	"MindsWallet/internal/errs"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	} else if errors.Is(err, errs.ErrValidationFailed) ||
		errors.Is(err, errs.ErrInvalidOperationType) ||
		errors.Is(err, errs.ErrUserAlreadyExists) ||
		errors.Is(err, errs.ErrNotEnoughBalance) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else if errors.Is(err, errs.ErrAccountNotFound) ||
		errors.Is(err, errs.ErrUserNotFound) ||
		errors.Is(err, errs.ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	} else if errors.Is(err, errs.ErrIncorrectUsernameOrPassword) ||
		errors.Is(err, errs.ErrUserIDNotFound) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
	} else if errors.Is(err, errs.ErrNoPermissionsToWithdraw) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("something went wrong: %s", err.Error()),
		})
	}
}
