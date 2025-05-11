package handlers

import (
	"net/http"
	"sport-app-backend/helper"
	"sport-app-backend/models"
	"sport-app-backend/services"

	"github.com/gin-gonic/gin"
)

type categoryFieldHandler struct {
	categoryFieldService services.CategoryFieldService
}

func NewCategoryFieldHandler(categoryFieldService services.CategoryFieldService) *categoryFieldHandler {
	return &categoryFieldHandler{categoryFieldService: categoryFieldService}
}

func (ch *categoryFieldHandler) CreateCategoryField(ctx *gin.Context) {
	input := &models.CategoryFieldRequest{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		errors := helper.FormatValidationError(err)
		errBindJson := helper.NewUnprocessableEntityError(errors[0])
		response := helper.APIResponse(errBindJson.Message(), errBindJson.Status(), "error", nil)
		ctx.JSON(errBindJson.Status(), response)
		return
	}

	categoryField, err := ch.categoryFieldService.CreateCategoryField(ctx, input)
	if err != nil {
		response := helper.APIResponse(err.Message(), err.Status(), "error", nil)
		ctx.JSON(err.Status(), response)
		return
	}

	formattedResponse := helper.APIResponse("Category field created successfully", http.StatusCreated, "success", categoryField)
	ctx.JSON(http.StatusCreated, formattedResponse)
}

func (ch *categoryFieldHandler) GetAllCategoryField(ctx *gin.Context) {
	categoryFields, err := ch.categoryFieldService.GetAllCategoryField(ctx)
	if err != nil {
		response := helper.APIResponse(err.Message(), err.Status(), "error", nil)
		ctx.JSON(err.Status(), response)
		return
	}

	formattedResponse := helper.APIResponse("Category fields found successfully", http.StatusOK, "success", categoryFields)
	ctx.JSON(http.StatusOK, formattedResponse)
}

func (ch *categoryFieldHandler) GetCategoryFieldByID(ctx *gin.Context) {
	id := ctx.Param("id")

	categoryField, err := ch.categoryFieldService.GetCategoryFieldByID(ctx, id)
	if err != nil {
		response := helper.APIResponse(err.Message(), err.Status(), "error", nil)
		ctx.JSON(err.Status(), response)
		return
	}

	response := helper.APIResponse("Category field found successfully", http.StatusOK, "success", categoryField)
	ctx.JSON(http.StatusOK, response)
}

func (ch *categoryFieldHandler) UpdateCategoryField(ctx *gin.Context) {
	id := ctx.Param("id")
	input := &models.CategoryFieldRequest{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		errors := helper.FormatValidationError(err)
		errBindJson := helper.NewUnprocessableEntityError(errors[0])
		response := helper.APIResponse(errBindJson.Message(), errBindJson.Status(), "error", nil)
		ctx.JSON(errBindJson.Status(), response)
		return
	}

	categoryField, err := ch.categoryFieldService.UpdateCategoryField(ctx, id, input)
	if err != nil {
		response := helper.APIResponse(err.Message(), err.Status(), "error", nil)
		ctx.JSON(err.Status(), response)
		return
	}

	formattedResponse := helper.APIResponse("Category field updated successfully", http.StatusOK, "success", categoryField)
	ctx.JSON(http.StatusOK, formattedResponse)
}

func (ch *categoryFieldHandler) DeleteCategoryField(ctx *gin.Context) {
	id := ctx.Param("id")

	categoryField, err := ch.categoryFieldService.DeleteCategoryField(ctx, id)
	if err != nil {
		response := helper.APIResponse(err.Message(), err.Status(), "error", nil)
		ctx.JSON(err.Status(), response)
		return
	}

	formattedResponse := helper.APIResponse("Category field deleted successfully", http.StatusOK, "success", categoryField)
	ctx.JSON(http.StatusOK, formattedResponse)
}
