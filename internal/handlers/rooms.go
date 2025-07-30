package handlers

import (
	"context"
	"errors"
	"fmt"
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

var (
	ErrRoomIDParamRequired = errors.New("room_id parameter is required")
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
	h.group.GET("/subscribe/:id", h.SubscribeToRoom)

}

func (h *roomsHttpHandler) CreateRoom(c *gin.Context) {
	roomId, err := h.controller.Create()
	if err != nil {
		respError(c, http.StatusInternalServerError, err.Error())
	}

	respSuccess(c, http.StatusCreated, roomId)
}

func (h *roomsHttpHandler) SubscribeToRoom(c *gin.Context) {
	roomId := c.Param("id")
	//todo: get userID from context / session
	userId := "add me"

	if roomId == "" {
		respError(c, http.StatusBadRequest, ErrRoomIDParamRequired.Error())
		fmt.Println("Err with nil roomid")
		return
	}

	err := h.controller.SubscribeToRoom(roomId, userId, c.Writer, c.Request)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			fmt.Printf("\n user id %v has disconnected from room id %v\n", userId, roomId)
			//TODO: damn so deep. Clean this up?
			err = h.controller.CleanUpEmptyRoomAfterDisconnect(roomId)
			if err != nil {
				respError(c, http.StatusInternalServerError, err.Error())
			}

			return
		}

		fmt.Println("Err with subscription ", err.Error())
		respError(c, http.StatusInternalServerError, err.Error())
	}
}

func (h *roomsHttpHandler) GetRooms(c *gin.Context) {
	// TODO: add
}
