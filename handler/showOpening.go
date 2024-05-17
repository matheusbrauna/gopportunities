package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusbrauna/gopportunities/schemas"
)

func ShowOpeningHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(c, http.StatusNotFound, "opening not found")
		return
	}

	sendSuccess(c, http.StatusOK, "show-opening", opening)
}
