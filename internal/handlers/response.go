package handlers

import (
	"github.com/gin-gonic/gin"
)

type HttpResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func respError(c *gin.Context, code int, message string) {
	h := &HttpResponse{
		Message: message,
	}

	c.JSON(code, h)
}

func respSuccess(c *gin.Context, code int, message string, data ...interface{}) {
	h := &HttpResponse{
		Message: message,
		Data:    data,
	}

	if len(data) > 0 {
		h.Data = data[0]
	}

	c.JSON(code, h)
}
