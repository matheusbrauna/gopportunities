package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}

func ListOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}

func CreateOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}

func DeleteOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}

func UpdateOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}
