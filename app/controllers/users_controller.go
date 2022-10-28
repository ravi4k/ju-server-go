package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"ju-go-server/app/models"
)

type UserController struct{}

func (UserController) GetAllUsers(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	users, err := models.User{}.FindAllUsers(limit)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (UserController) GetUser(ctx *gin.Context) {
	userId := ctx.Param("userId")
	user, err := models.User{}.FindUser(userId)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
