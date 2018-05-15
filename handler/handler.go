package handler

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

func Get(c *gin.Context) {
	fmt.Println("get")
	c.String(http.StatusOK, "get")
}
func Post(c *gin.Context) {
	fmt.Println("post")
	c.String(http.StatusOK, "post")
}
func Put(c *gin.Context) {
	fmt.Println("put")
	c.String(http.StatusOK, "put")
}
func Delete(c *gin.Context) {
	fmt.Println("delete")
	c.String(http.StatusOK, "delete")
}
func Patch(c *gin.Context) {
	fmt.Println("patch")
	c.String(http.StatusOK, "patch")
}
func Options(c *gin.Context) {
	fmt.Println("options")
}

func Head(c *gin.Context) {
	fmt.Println("head")
}
