package main

import (
	"github.com/ProgressiveParanoia/go-game-server/internal/controller"
	"github.com/ProgressiveParanoia/go-game-server/internal/handlers"
	"github.com/ProgressiveParanoia/go-game-server/internal/repo/memory"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the router
	r := gin.Default()

	memRepo := memory.NewMemoryRepository()

	// Initialize the user controller
	userController := controller.NewUserController(memRepo)

	handlers.InitRouter(r, userController)
}
