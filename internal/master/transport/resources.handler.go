package transport

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/harleywinston/x-manager/internal/master/consts"
	"github.com/harleywinston/x-manager/internal/master/models"
	"github.com/harleywinston/x-manager/internal/master/services"
)

type ResourcesHandler struct {
	resourcesService services.ResourcesService
}

func (h *ResourcesHandler) AddResourcesHandler(ctx *gin.Context) {
	var resource models.Resources
	err := ctx.BindJSON(&resource)
	if err != nil {
		ctx.JSON(consts.BIND_JSON_ERROR.Code, gin.H{
			"message": consts.BIND_JSON_ERROR.Message,
			"detail":  err.Error(),
		})
		return
	}

	err = h.resourcesService.AddResourcesService(resource)
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
		"detail":  fmt.Sprintf(`Resource serverIP: %s`, resource.ServerIp),
	})
}

func (h *ResourcesHandler) GetResourcesHandler(ctx *gin.Context) {
	var resource models.Resources
	err := ctx.BindJSON(&resource)
	if err != nil {
		ctx.JSON(consts.BIND_JSON_ERROR.Code, gin.H{
			"message": consts.BIND_JSON_ERROR.Message,
			"detail":  err.Error(),
		})
		return
	}

	res, err := h.resourcesService.GetResourcesService(resource)
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

func (h *ResourcesHandler) DeleteResourcesHandler(ctx *gin.Context) {
	var resource models.Resources
	err := ctx.BindJSON(&resource)
	if err != nil {
		ctx.JSON(consts.BIND_JSON_ERROR.Code, gin.H{
			"message": consts.BIND_JSON_ERROR.Message,
			"detail":  err.Error(),
		})
		return
	}

	err = h.resourcesService.DeleteResourcesService(resource)
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
		"detail":  fmt.Sprintf(`Resource serverIP: %s`, resource.ServerIp),
	})
}
