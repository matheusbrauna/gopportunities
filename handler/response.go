package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func sendError(c *gin.Context, code int, message string) {
	c.Header("Content-type", "application/json")
	c.JSON(code, gin.H{
		"message":  message,
		"erroCode": code,
	})
}

func sendSuccess(c *gin.Context, code int, operation string, data interface{}) {
	c.Header("Content-type", "application/json")
	c.JSON(code, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", operation),
		"data":    data,
	})
}
