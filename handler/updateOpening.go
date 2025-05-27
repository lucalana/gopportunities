package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucalana/gopportunities/schemas"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error %v", err.Error())
		sendError(ctx, 404, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening).Error; err != nil {
		sendError(ctx, http.StatusBadRequest, "Opening not found")
	}
	// Update opening
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}
	// Save opening
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("Error update opening %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
	}

	sendSuccess(ctx, "update-opening", opening)
}
