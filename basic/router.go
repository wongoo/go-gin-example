package main

import "github.com/gin-gonic/gin"

var DB = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	gin.DisableConsoleColor()

	// 创建带有默认中间件的路由:
	// 日志与恢复中间件
	router := gin.Default()

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	//---------------------------------------------
	// Get user value from query parameter
	router.GET("/user", func(c *gin.Context) {
		user := c.Query("name")
		value, ok := DB[user]
		if ok {
			c.JSON(200, gin.H{"user": user, "value": value})
		} else {
			// 从请求参数获得value值，并设置默认值
			value = c.DefaultQuery("value", "default")
			c.JSON(200, gin.H{"user": user, "status": value})
		}
	})

	//---------------------------------------------
	// Get user value from path variable
	router.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := DB[user]
		if ok {
			c.JSON(200, gin.H{"user": user, "value": value})
		} else {
			c.JSON(200, gin.H{"user": user, "status": "no value"})
		}
	})

	//---------------------------------------------
	//Authorized group (uses gin.BasicAuth() middleware)
	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	//---------------------------------------------
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		// Bind using query & form & post parameter
		if c.Bind(&json) == nil {
			DB[user] = json.Value
			c.JSON(200, gin.H{"status": "ok"})
		}
	})

	return router
}
