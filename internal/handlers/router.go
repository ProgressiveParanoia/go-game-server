package handlers

import (
	"github.com/ProgressiveParanoia/go-game-server/internal/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, userController *controller.User) {
	r.Use(recoveryMiddleware())

	userGroup := r.Group("/users")

	userGroup.Use(recoveryMiddleware())

	userHandler := NewUserHttpHandler(userGroup, userController)

	userHandler.RegisterRoutes()

	r.Run() // listen and serve on
}
