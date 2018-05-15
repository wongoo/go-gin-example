package main

import (
	"github.com/thinkerou/favicon"
	"net/http"
	"time"
	"os"
	"os/signal"
	"log"
	"context"
)

func main() {
	router := setupRouter()
	router.Use(favicon.New("./favicon.ico"))

	//---------------------------------------------
	// Listen and Server in 0.0.0.0:8080
	// 默认服务器
	//router.Run(":8080")

	//---------------------------------------------
	// HTTP 服务器
	// http.ListenAndServe(":8080", router)

	//---------------------------------------------
	// 自定义 HTTP 服务器的配置
	srv := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
