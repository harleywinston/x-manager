package transport

import (
	"github.com/gin-gonic/gin"

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
		ctx.JSON(500, gin.H{
			"message": "bad request",
			"error":   err,
		})
		return
	}

	err = h.resourcesService.AddResourcesService(resource)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "bad request",
			"error":   err,
		})
		return
	}
}

func (h *ResourcesHandler) GetResourcesHandler(ctx *gin.Context) {
	var resource models.Resources
	err := ctx.BindJSON(&resource)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "bad request",
			"error":   err,
		})
		return
	}

	res, err := h.resourcesService.GetResourcesService(resource)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "bad request",
			"error":   err,
		})
		return
	}
	ctx.JSON(200, res)
}

func (h *ResourcesHandler) DeleteResourcesHandler(ctx *gin.Context) {
	var resource models.Resources
	err := ctx.BindJSON(&resource)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "bad request",
			"error":   err,
		})
		return
	}

	err = h.resourcesService.DeleteResourcesService(resource)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "bad request",
			"error":   err,
		})
		return
	}
}
