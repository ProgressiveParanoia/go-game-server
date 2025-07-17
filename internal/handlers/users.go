package handlers

import (
	"fmt"
	"net/http"

	"github.com/ProgressiveParanoia/go-game-server/internal/controller"
	"github.com/gin-gonic/gin"
)

type (
	HttpNewUserPost struct {
		Name           string `json:"name"`
		DeviceID       string `json:"device_id"`
		ProfilePicture string `json:"profile_picture"`
	}
	userHttpHandler struct {
		group      *gin.RouterGroup
		controller *controller.User
	}
)

func NewUserHttpHandler(group *gin.RouterGroup, controller *controller.User) *userHttpHandler {
	return &userHttpHandler{
		group:      group,
		controller: controller,
	}
}

func (h *userHttpHandler) RegisterRoutes() {
	h.group.GET("/:id", h.GetUserByID)
	h.group.GET("/", h.GetAllUsers)
	h.group.POST("/create", h.CreateUser)
}

func (h *userHttpHandler) GetAllUsers(c *gin.Context) {
	users, err := h.controller.GetAll()
	if err != nil {
		respError(c, http.StatusInternalServerError, "failed to retrieve users")
		return
	}

	respSuccess(c, http.StatusOK, "users retrieved successfully", users)

}

func (h *userHttpHandler) GetUserByID(c *gin.Context) {
	user, err := h.controller.GetByID(c.Param("id"))
	if err != nil {
		respError(c, http.StatusNotFound, "user not found")
		return
	}

	respSuccess(c, http.StatusOK, "user retrieved successfully", user)
}

func (h *userHttpHandler) CreateUser(c *gin.Context) {
	body := HttpNewUserPost{}
	if err := c.Bind(&body); err != nil {
		respError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	newUserID, err := h.controller.Create(body.Name, body.DeviceID, body.ProfilePicture)
	if err != nil {
		if err == controller.ErrUserAlreadyExists {
			respError(c, http.StatusConflict, "user already exists")
			return
		} else {
			respError(c, http.StatusInternalServerError,
				fmt.Sprintf("unexpected error failed to create user with device ID %v", body.DeviceID))
			return
		}
	}

	type httpNewUserPostResponse struct {
		ID string `json:"id"`
	}

	respSuccess(c, http.StatusOK, "user successfully created",
		httpNewUserPostResponse{ID: newUserID})
}
