package handlers

import (
	"github.com/ProgressiveParanoia/go-game-server/internal/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, userController *controller.User) {
	//TODO: Add authorization middleware.
	// Anyone can access the API right now.
	userGroup := r.Group("/users")

	userHandler := NewUserHttpHandler(userGroup, userController)

	userHandler.RegisterRoutes()

	r.Run() // listen and serve on
}
