package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucalana/gopportunities/schemas"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error listing openings")
	}

	sendSuccess(ctx, "list-openings", openings)
}
