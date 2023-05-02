package app

import (
	"github.com/gin-gonic/gin"

	"github.com/harleywinston/x-manager/internal/master/transport"
)

func InitApp() error {
	r := gin.Default()
	r.GET("/user", transport.GetUserHandler)
	r.POST("/user", transport.AddUserHandler)
	r.DELETE("/user", transport.DeleteUserHandler)

	return r.Run()
}
