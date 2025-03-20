package handlers

import (
	"net/http"
	"sport-app-backend/helper"
	"sport-app-backend/models"
	"sport-app-backend/services"

	"github.com/gin-gonic/gin"
)

type userOwnerHandler struct {
	userOwnerService services.UserOwnerService
}

func NewUserOwnerHandler(userOwnerService services.UserOwnerService) *userOwnerHandler {
	return &userOwnerHandler{userOwnerService: userOwnerService}
}

func (uoh *userOwnerHandler) CreateUserOwner(ctx *gin.Context) {
	userOwnerPayLoad := &models.RegisterUserOwnerRequest{}

	if err := ctx.ShouldBindJSON(userOwnerPayLoad); err != nil {
		errors := helper.FormatValidationError(err)

		errBindJson := helper.NewUnprocessableEntityError(errors[0])
		response := helper.APIResponse(errBindJson.Message(), errBindJson.Status(), "error", nil)
		ctx.JSON(errBindJson.Status(), response)
		return
	}

	userOwner, err := uoh.userOwnerService.CreateUserOwner(ctx, userOwnerPayLoad)
	if err != nil {
		response := helper.APIResponse(err.Message(), err.Status(), "error", nil)
		ctx.JSON(err.Status(), response)
		return
	}

	formattedResponse := helper.APIResponse("User owner created successfully", http.StatusCreated, "success", userOwner)
	ctx.JSON(http.StatusCreated, formattedResponse)
}
