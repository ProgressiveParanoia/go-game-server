package handlers

import (
	"github.com/ProgressiveParanoia/go-game-server/internal/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, userController *controller.User, roomController *controller.Room) {
	r.Use(recoveryMiddleware())

	userGroup := r.Group("/users")
	roomsGroup := r.Group("/rooms")

	userGroup.Use(recoveryMiddleware())

	userHandler := NewUserHttpHandler(userGroup, userController)
	roomHandler := NewRoomHttpHandler(roomsGroup, roomController)

	userHandler.RegisterRoutes()
	roomHandler.RegisterRoutes()

	r.Run() // listen and serve on
}
