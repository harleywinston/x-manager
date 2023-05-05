package transport

import (
	"github.com/gin-gonic/gin"

	"github.com/harleywinston/x-manager/internal/master/models"
	"github.com/harleywinston/x-manager/internal/master/services"
)

type GroupsHandler struct {
	service services.GroupsService
}

func (h *GroupsHandler) AddGroupsHandler(ctx *gin.Context) {
	var group models.Groups
	if err := ctx.BindJSON(&group); err != nil {
		ctx.JSON(500, gin.H{
			"message": "Bad Request!",
			"error":   err.Error(),
		})
		return
	}

	if err := h.service.AddGroupsService(group); err != nil {
		ctx.JSON(500, gin.H{
			"message": "Interal error!",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Group added!",
	})
}

func (h *GroupsHandler) GetGroupsHandler(ctx *gin.Context) {
	var group models.Groups
	if err := ctx.BindJSON(&group); err != nil {
		ctx.JSON(500, gin.H{
			"message": "Bad Request!",
			"error":   err.Error(),
		})
		return
	}

	res, err := h.service.GetGroupsService(group)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Interal error!",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"group": res,
	})
}

func (h *GroupsHandler) DeleteGroupsHandler(ctx *gin.Context) {
	var group models.Groups
	if err := ctx.BindJSON(&group); err != nil {
		ctx.JSON(500, gin.H{
			"message": "Bad Request!",
			"error":   err.Error(),
		})
		return
	}

	if err := h.service.DeleteGroupsService(group); err != nil {
		ctx.JSON(500, gin.H{
			"message": "Interal error!",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Group deleted!",
	})
}
