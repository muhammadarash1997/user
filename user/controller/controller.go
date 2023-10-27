package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadarash1997/user/user/model"
	"github.com/muhammadarash1997/user/user/service"
)

type controller struct {
	service service.Service
}

func NewController(service service.Service) *controller {
	return &controller{
		service: service,
	}
}

func (c *controller) RegisterUserHandler(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = c.service.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Password = ""
	ctx.JSON(http.StatusCreated, user)
}

func (c *controller) GetUserHandler(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(model.User)
	user, err := c.service.FindUser(currentUser.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
