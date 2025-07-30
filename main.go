package main

import (
	"fmt"
	"os"

	"github.com/ProgressiveParanoia/go-game-server/internal/controller"
	"github.com/ProgressiveParanoia/go-game-server/internal/handlers"
	"github.com/ProgressiveParanoia/go-game-server/internal/repo/memory"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		fmt.Errorf("error running application %v", err)
		os.Exit(1)
	}
}

func run() error {
	// Initialize the router
	r := gin.Default()

	userRepo := memory.NewUserRepository()

	// Initialize controllers
	userController := controller.NewUserController(userRepo)
	roomController := controller.NewRoomController(nil)

	handlers.InitRouter(r, userController, roomController)

	return nil
}
