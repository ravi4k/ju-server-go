package server

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	setupMiddlewares(router)
	return router
}

func setupMiddlewares(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}
