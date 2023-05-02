package master

import (
	"github.com/gin-gonic/gin"

	"github.com/harleywinston/x-manager/internal/master/transport"
)

func InitApp() error {
	r := gin.Default()

	usersHandlers := transport.UsersHandler{}
	r.GET("/user", usersHandlers.GetUserHandler)
	r.POST("/user", usersHandlers.AddUserHandler)
	r.DELETE("/user", usersHandlers.DeleteUserHandler)

	return r.Run()
}
