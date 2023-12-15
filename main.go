package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go-exercise/api"
	configs "go-exercise/configs"

	"github.com/gin-gonic/gin"
)

func initWebServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.ForceConsoleColor()

	router := gin.Default()

	// Route
	api.RegisterRoutes(router)

	return router
}

func main() {
	router := initWebServer()

	fmt.Println("Server is listening and serving HTTP on", configs.ServerPort)
	server := &http.Server{
		Addr:    configs.ServerPort,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %v\n", err)
			os.Exit(1)
		}
	}()

	// Set up a channel for graceful shutdown
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, os.Interrupt)

	// Block until a signal is received
	<-quitChannel
	fmt.Println("The server is shutting down...")

	// Create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt a graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down server: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Exit program")
	os.Exit(0)
}
