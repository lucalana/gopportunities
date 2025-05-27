package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucalana/gopportunities/schemas"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	// Find opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "Opening not found")
		return
	}
	// Delete opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error on deleting opening")
		return
	}
	sendSuccess(ctx, "Delete opening", opening)
}
