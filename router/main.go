package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wongoo/go-gin-example/handler"
)

func main() {
	router := gin.Default();
	router.GET("/get", handler.Get)
	router.POST("/post", handler.Post)
	router.PUT("/put", handler.Put)
	router.DELETE("/delete", handler.Delete)
	router.PATCH("/patch", handler.Patch)
	router.HEAD("/head", handler.Head)
	router.OPTIONS("/options", handler.Options)

	router.Run(":8081")
}
