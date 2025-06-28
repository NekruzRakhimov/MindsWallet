package controller

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RunServer() *gin.Engine {
	router := gin.Default()

	router.GET("/", Ping)

	authG := router.Group("/auth")
	{
		authG.POST("/sign-up", SignUp)
		authG.POST("/sign-in", SignIn)
	}

	apiG := router.Group("/api", checkUserAuthentication)

	accountsG := apiG.Group("/accounts")
	{
		accountsG.GET("", GetAllAccounts)
		accountsG.GET("/:id", GetAccountByID)
		accountsG.PATCH("/balance/:id", UpdateAccountBalance)
	}

	profileG := apiG.Group("/profile")
	{
		profileG.GET("")
		profileG.PUT("")
	}

	//if err := router.Run(configs.AppSettings.AppParams.PortRun); err != nil {
	//	logger.Error.Printf("[controller] RunServer():  Error during running HTTP server: %s", err.Error())
	//	return err
	//}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
