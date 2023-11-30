package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Response(c *gin.Context, status int, data any) {
	c.JSON(status, data)
}

func Success(c *gin.Context, status int, body any) {
	Response(c, status, body)
}

// Error creates a new error with the given status code and the message
// formatted according to args and format.
func Error(c *gin.Context, status int, format string, args ...any) {
	err := errorResponse{
		Code:    strings.ReplaceAll(strings.ToLower(http.StatusText(status)), " ", "_"),
		Message: fmt.Sprintf(format, args...),
		Status:  status,
	}

	Response(c, status, err)
}

// InternalServerError creates a default InternalServerError response
func InternalServerError(ctx *gin.Context) {
	Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
}