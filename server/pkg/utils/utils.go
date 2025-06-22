package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type HTTPResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func CreateSuccessfulHTTPResponse(message string, data any) HTTPResponse {
	return HTTPResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func CreateErrorHTTPResponse(message string, err error) HTTPResponse {
	if err != nil {
		message = message + err.Error()
	}

	return HTTPResponse{
		Success: false,
		Error:   message,
	}
}

func GetUserIdFromContext(c *gin.Context) (int64, error) {
	val, exists := c.Get("user_id")
	if !exists {
		return 0, fmt.Errorf("user_id not found in context")
	}

	user_id, ok := (val.(float64))
	if !ok {
		return 0, fmt.Errorf("user_id has unexpected type")
	}

	return int64(user_id), nil
}
