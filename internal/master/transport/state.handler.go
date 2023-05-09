package transport

import (
	"github.com/gin-gonic/gin"

	"github.com/harleywinston/x-manager/internal/master/consts"
	"github.com/harleywinston/x-manager/internal/master/services"
)

type StateHandlers struct {
	stateServices services.StateServices
}

func (h *StateHandlers) GetStateHandler(ctx *gin.Context) {
	serverIP := ctx.ClientIP()
	state, err := h.stateServices.GetStateService(serverIP)
	if err != nil {
		if e, ok := err.(*consts.CustomError); ok {
			ctx.JSON(e.Code, gin.H{
				"message": e.Message,
				"detail":  e.Detail,
			})
		}
		return
	}

	ctx.JSON(200, state)
}
