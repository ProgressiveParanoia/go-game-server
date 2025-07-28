package handlers

import (
	"net/http"

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
	h.group.GET("/", h.GetRooms)
	h.group.POST("/create", h.CreateRoom)
	h.group.PUT("/join/:id", h.JoinRoom)
	h.group.GET("/subscribe/:id", h.SubscribeToRoom)

}

func (h *roomsHttpHandler) CreateRoom(c *gin.Context) {
	//does room exist?
}

func (h *roomsHttpHandler) JoinRoom(c *gin.Context) {
	//add myself to session, then create a socket?

	//does the room exist?
	//get associated user data from session
	//append the user data to the ongoing room
	//create socket for client to connect to
}

func (h *roomsHttpHandler) SubscribeToRoom(c *gin.Context) {
	err := h.controller.SubscribeToRoom(c.Param("room_id"), c.Param("user_id"), c.Writer, c.Request)
	if err != nil {
		respError(c, http.StatusInternalServerError, err.Error())
	}
}

func (h *roomsHttpHandler) GetRooms(c *gin.Context) {

}
