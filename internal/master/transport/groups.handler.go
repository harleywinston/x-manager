package transport

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/harleywinston/x-manager/internal/master/consts"
	"github.com/harleywinston/x-manager/internal/master/models"
	"github.com/harleywinston/x-manager/internal/master/services"
)

type GroupsHandler struct {
	service services.GroupsService
}

func (h *GroupsHandler) AddGroupsHandler(ctx *gin.Context) {
	var group models.Groups
	err := ctx.BindJSON(&group)
	if err != nil {
		ctx.JSON(consts.BIND_JSON_ERROR.Code, gin.H{
			"message": consts.BIND_JSON_ERROR.Message,
			"detail":  err.Error(),
		})
		return
	}

	err = h.service.AddGroupsService(group)
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
		"detail":  fmt.Sprintf(`Group resourceID: %d`, group.ResourcesID),
	})
}

func (h *GroupsHandler) GetGroupsHandler(ctx *gin.Context) {
	var group models.Groups
	err := ctx.BindJSON(&group)
	if err != nil {
		ctx.JSON(consts.BIND_JSON_ERROR.Code, gin.H{
			"message": consts.BIND_JSON_ERROR.Message,
			"detail":  err.Error(),
		})
		return
	}

	res, err := h.service.GetGroupsService(group)
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

func (h *GroupsHandler) DeleteGroupsHandler(ctx *gin.Context) {
	var group models.Groups
	err := ctx.BindJSON(&group)
	if err != nil {
		ctx.JSON(consts.BIND_JSON_ERROR.Code, gin.H{
			"message": consts.BIND_JSON_ERROR.Message,
			"detail":  err.Error(),
		})
		return
	}

	err = h.service.DeleteGroupsService(group)
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
		"detail":  fmt.Sprintf(`Group resourceID: %d`, group.ResourcesID),
	})
}
