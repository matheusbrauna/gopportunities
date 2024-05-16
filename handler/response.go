package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(c *gin.Context, code int, message string) {
	c.Header("Content-type", "application/json")
	c.JSON(code, gin.H{
		"message":  message,
		"erroCode": code,
	})
}

func sendSuccess(c *gin.Context, operation string, data interface{}) {
	c.Header("Content-type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s", operation),
		"data":    data,
	})
}