package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code":    "aixiba",
			"message": "justtest",
		})
	})
	if err := router.Run(); err != nil {
		panic(err)
	}
}
