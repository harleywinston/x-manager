package master

import "github.com/gin-gonic/gin"

func SetupMaster() error {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r.Run()
}
