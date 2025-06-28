package controller

import (
	"MindsWallet/internal/models"
	"MindsWallet/internal/service"
	"MindsWallet/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SignUp godoc
// @Summary Регистрация пользователя
// @Description Создаёт нового пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "Данные пользователя"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /auth/sign-up [post]
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

// SignIn godoc
// @Summary Авторизация пользователя
// @Description Вход по логину и паролю
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.UserSignIn true "Учётные данные"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /auth/sign-in [post]
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
