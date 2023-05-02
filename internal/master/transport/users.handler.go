package transport

import (
	"github.com/gin-gonic/gin"

	"github.com/harleywinston/x-manager/internal/master/models"
	"github.com/harleywinston/x-manager/internal/master/services"
)

type UsersHandler struct {
	userService services.UsersService
}

func (h *UsersHandler) GetUserHandler(ctx *gin.Context) {
	var user models.UsersModel
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "bad request",
			"error":   err,
		})
		return
	}
	if user.Email == "" {
		ctx.JSON(500, gin.H{
			"message": "bad request",
			"error":   "Email not provided!",
		})
		return
	}

	res, err := h.userService.GetUserService(user)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "server error",
			"error":   err,
		})
		return
	}

	ctx.JSON(200, res)
}

func (h *UsersHandler) AddUserHandler(ctx *gin.Context) {
	var user models.UsersModel
	err := ctx.Bind(&user)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "bad request",
			"error":   err,
		})
		return
	}
	if user.Email == "" {
		ctx.JSON(500, gin.H{
			"message": "bad request",
			"error":   "Email not provided!",
		})
		return
	}
	if user.Username == "" {
		ctx.JSON(500, gin.H{
			"message": "bad request",
			"error":   "Username not provided!",
		})
		return
	}

	err = h.userService.AddUserService(user)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "server error",
			"error":   err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "user added!",
	})
}

func (h *UsersHandler) DeleteUserHandler(ctx *gin.Context) {
	user := models.UsersModel{}
	err := ctx.Bind(&user)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "bad request",
			"error":   err,
		})
		return
	}
	err = h.userService.DeleteUserService(user)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "server error",
			"error":   err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "user deleted!",
	})
}
