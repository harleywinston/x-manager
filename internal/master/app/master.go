package app

import "github.com/gin-gonic/gin"

func InitApp() error {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "fuck you",
		})
	})

	return r.Run()
}
