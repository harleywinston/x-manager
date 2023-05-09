package transport

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/harleywinston/x-manager/internal/master/consts"
	"github.com/harleywinston/x-manager/internal/master/models"
	"github.com/harleywinston/x-manager/internal/master/services"
)

type UsersHandler struct {
	userService services.UsersService
}

func (h *UsersHandler) GetUserHandler(ctx *gin.Context) {
	var user models.Users
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(consts.BIND_JSON_ERROR.Code, gin.H{
			"message": consts.BIND_JSON_ERROR.Message,
			"detail":  err.Error(),
		})
		return
	}

	res, err := h.userService.GetUserService(user)
	if err != nil {
		if e, ok := err.(*consts.CustomError); ok {
			ctx.JSON(e.Code, gin.H{
				"message": e.Message,
				"detail":  e.Detail,
			})
		}
		return
	}

	ctx.JSON(200, res)
}

func (h *UsersHandler) GetUserConfigs(ctx *gin.Context) {
	var user models.Users
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(consts.BIND_JSON_ERROR.Code, gin.H{
			"message": consts.BIND_JSON_ERROR.Message,
			"detail":  err.Error(),
		})
		return
	}

	configsStr, err := h.userService.GetUserConfigs(user)
	if err != nil {
		if e, ok := err.(*consts.CustomError); ok {
			ctx.JSON(e.Code, gin.H{
				"message": e.Message,
				"detail":  e.Detail,
			})
		}
		return
	}

	ctx.Writer.Header().Set("Content-Type", "text/plain")
	ctx.String(200, configsStr)
}

func (h *UsersHandler) AddUserHandler(ctx *gin.Context) {
	var user models.Users
	err := ctx.Bind(&user)
	if err != nil {
		ctx.JSON(consts.BIND_JSON_ERROR.Code, gin.H{
			"message": consts.BIND_JSON_ERROR.Message,
			"detail":  err.Error(),
		})
		return
	}

	err = h.userService.AddUserService(user)
	if err != nil {
		if e, ok := err.(*consts.CustomError); ok {
			ctx.JSON(e.Code, gin.H{
				"message": e.Message,
				"detail":  e.Detail,
			})
		}
		return
	}

	ctx.JSON(consts.ADD_SUCCESS.Code, gin.H{
		"message": consts.ADD_SUCCESS.Message,
		"detail":  fmt.Sprintf(`User email: %s`, user.Email),
	})
}

func (h *UsersHandler) DeleteUserHandler(ctx *gin.Context) {
	user := models.Users{}
	err := ctx.Bind(&user)
	if err != nil {
		ctx.JSON(consts.BIND_JSON_ERROR.Code, gin.H{
			"message": consts.BIND_JSON_ERROR.Message,
			"detail":  err.Error(),
		})
		return
	}

	err = h.userService.DeleteUserService(user)
	if err != nil {
		if e, ok := err.(*consts.CustomError); ok {
			ctx.JSON(e.Code, gin.H{
				"message": e.Message,
				"detail":  e.Detail,
			})
		}
		return
	}

	ctx.JSON(consts.DELETE_SUCCESS.Code, gin.H{
		"message": consts.DELETE_SUCCESS.Message,
		"detail":  fmt.Sprintf(`User email: %s`, user.Email),
	})
}
