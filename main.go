package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/wongoo/go-gin-example/config"
)

func main() {

	router := gin.Default()

	router.GET("/env", func(context *gin.Context) {
		context.String(http.StatusOK, "current env is %v", config.GIN_GO_ENV)
	})

	router.Run(":8081")
}
