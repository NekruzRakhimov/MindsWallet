package controller

import (
	"MindsWallet/internal/models"
	"MindsWallet/internal/service"
	"MindsWallet/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(c *gin.Context) {
	var u models.User

	if err := c.ShouldBindJSON(&u); err != nil {
		HandleError(c, err)
		return
	}

	if err := service.CreateUser(u); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}

func SignIn(c *gin.Context) {
	// получить идентификатор и пароль
	var u models.UserSignIn
	if err := c.ShouldBindJSON(&u); err != nil {
		HandleError(c, err)
		return
	}

	// отправить в бд запрос на проверку есть ли такой
	// пользователь с таким паролем
	user, err := service.GetUserByUsernameAndPassword(u.Username, u.Password)
	if err != nil {
		HandleError(c, err)

		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})
}
