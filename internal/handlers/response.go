package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func respError(c *gin.Context, code int, message string) {
	h := &HttpResponse{
		Code:    code,
		Message: message,
	}

	c.JSON(http.StatusInternalServerError, h)
}

func respSuccess(c *gin.Context, code int, message string, data ...interface{}) {
	h := &HttpResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}

	if len(data) > 0 {
		h.Data = data[0]
	}

	c.JSON(code, h)
}
