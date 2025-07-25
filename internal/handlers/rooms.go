package handlers

import (
	"github.com/ProgressiveParanoia/go-game-server/internal/controller"
	"github.com/gin-gonic/gin"
)

type (
	roomsHttpHandler struct {
		group      *gin.RouterGroup
		controller *controller.Room
	}
)

func NewRoomHttpHandler(group *gin.RouterGroup, controller *controller.Room) *roomsHttpHandler {
	return &roomsHttpHandler{
		group:      group,
		controller: controller,
	}
}

func (h *roomsHttpHandler) RegisterRoutes() {

}
