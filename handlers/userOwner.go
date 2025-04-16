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

func (uoh *userOwnerHandler) LoginUserOwner(ctx *gin.Context) {
	userOwnerPayLoad := &models.LoginUserOwnerRequest{}

	if err := ctx.ShouldBindJSON(userOwnerPayLoad); err != nil {
		errors := helper.FormatValidationError(err)

		errBindJson := helper.NewUnprocessableEntityError(errors[0])
		response := helper.APIResponse(errBindJson.Message(), errBindJson.Status(), "error", nil)
		ctx.JSON(errBindJson.Status(), response)
		return
	}

	userOwner, err := uoh.userOwnerService.LoginUserOwner(ctx, userOwnerPayLoad)
	if err != nil {
		response := helper.APIResponse(err.Message(), err.Status(), "error", nil)
		ctx.JSON(err.Status(), response)
		return
	}

	formattedResponse := helper.APIResponse("User owner logged in successfully", http.StatusOK, "success", userOwner)
	ctx.JSON(http.StatusOK, formattedResponse)
}

func (uoh *userOwnerHandler) RequestResetPasswordHandler(ctx *gin.Context) {
	var request models.ResetPasswordRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		errors := helper.FormatValidationError(err)
		errResp := helper.NewUnprocessableEntityError(errors[0])
		ctx.JSON(errResp.Status(), helper.APIResponse(errResp.Message(), errResp.Status(), "error", nil))
		return
	}

	// Panggil service untuk request reset password
	err := uoh.userOwnerService.RequestResetPassword(ctx, request.Email)
	if err != nil {
		ctx.JSON(err.Status(), helper.APIResponse(err.Message(), err.Status(), "error", nil))
		return
	}

	ctx.JSON(http.StatusOK, helper.APIResponse("OTP has been sent to your email", http.StatusOK, "success", nil))
}

// Reset Password (Verifikasi OTP dan Ganti Password)
func (uoh *userOwnerHandler) ResetPasswordHandler(ctx *gin.Context) {
	var request models.ResetPasswordWithOTPRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		errors := helper.FormatValidationError(err)
		errResp := helper.NewUnprocessableEntityError(errors[0])
		ctx.JSON(errResp.Status(), helper.APIResponse(errResp.Message(), errResp.Status(), "error", nil))
		return
	}

	// Panggil service untuk verifikasi OTP dan reset password
	err := uoh.userOwnerService.ResetPassword(ctx, request.Email, request.OTP, request.NewPassword)
	if err != nil {
		ctx.JSON(err.Status(), helper.APIResponse(err.Message(), err.Status(), "error", nil))
		return
	}

	ctx.JSON(http.StatusOK, helper.APIResponse("Password has been reset successfully", http.StatusOK, "success", nil))
}