package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusbrauna/gopportunities/schemas"
)

func ListOpeningHandler(c *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error listing openings")
	}

	sendSuccess(c, http.StatusOK, "list-openings", openings)
}
