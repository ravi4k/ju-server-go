package config

import (
	"github.com/gin-gonic/gin"
	"ju-go-server/app/controllers"
)

func SetupRoutes(router gin.IRouter) {
	setupUserRoutes(router)
}

func setupUserRoutes(router gin.IRouter) {
	userController := controllers.UserController{}
	userRouter := router.Group("/users")

	userRouter.GET("/", userController.GetAllUsers)
	userRouter.GET("/:userId", userController.GetUser)
}
