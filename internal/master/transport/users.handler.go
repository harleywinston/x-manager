package transport

import (
	"github.com/gin-gonic/gin"

	"github.com/harleywinston/x-manager/internal/master/models"
	"github.com/harleywinston/x-manager/internal/master/services"
)

type UsersHandler struct {
	userSerivce services.UsersService
}

func (h *UsersHandler) GetUserHandler(ctx *gin.Context) {
	user := models.UsersModel{}
	user.Email = ""
	_, _ = h.userSerivce.GetUserService(user)
}

func (h *UsersHandler) AddUserHandler(ctx *gin.Context) {
	user := models.UsersModel{}
	user.Email = ""
	_ = h.userSerivce.AddUserService(user)
}

func (h *UsersHandler) DeleteUserHandler(ctx *gin.Context) {
	user := models.UsersModel{}
	user.Email = ""
	_ = h.userSerivce.DeleteUserService(user)
}
