package handlers

import (
	"net/http"
	"sport-app-backend/helper"
	"sport-app-backend/models"
	"sport-app-backend/services"

	"github.com/gin-gonic/gin"
)

type categoryProductHandler struct {
	categoryProductService services.CategoryProductService
}

func NewCategoryProductHandler(categoryProductService services.CategoryProductService) *categoryProductHandler {
	return &categoryProductHandler{categoryProductService: categoryProductService}
}

func (cph *categoryProductHandler) CreateCategoryProduct(ctx *gin.Context) {
	input := &models.CategoryProductRequest{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		errors := helper.FormatValidationError(err)
		errBindJson := helper.NewUnprocessableEntityError(errors[0])
		response := helper.APIResponse(errBindJson.Message(), errBindJson.Status(), "error", nil)
		ctx.JSON(errBindJson.Status(), response)
		return
	}

	categoryProduct, err := cph.categoryProductService.CreateCategoryProduct(ctx, input)
	if err != nil {
		response := helper.APIResponse(err.Message(), err.Status(), "error", nil)
		ctx.JSON(err.Status(), response)
		return
	}

	formattedResponse := helper.APIResponse("Category product created successfully", http.StatusCreated, "success", categoryProduct)
	ctx.JSON(http.StatusCreated, formattedResponse)
}

func (cph *categoryProductHandler) GetCategoryProductByID(ctx *gin.Context) {
	id := ctx.Param("id")

	categoryProduct, err := cph.categoryProductService.GetCategoryProductByID(ctx, id)
	if err != nil {
		response := helper.APIResponse(err.Message(), err.Status(), "error", nil)
		ctx.JSON(err.Status(), response)
		return
	}

	response := helper.APIResponse("Category product found successfully", http.StatusOK, "success", categoryProduct)
	ctx.JSON(http.StatusOK, response)
}

func (cph *categoryProductHandler) GetAllCategoryProduct(ctx *gin.Context) {
	categoryProducts, err := cph.categoryProductService.GetAllCategoryProduct(ctx)
	if err != nil {
		response := helper.APIResponse(err.Message(), err.Status(), "error", nil)
		ctx.JSON(err.Status(), response)
		return
	}

	formattedResponse := helper.APIResponse("Category products found successfully", http.StatusOK, "success", categoryProducts)
	ctx.JSON(http.StatusOK, formattedResponse)
}

func (cph *categoryProductHandler) UpdateCategoryProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	input := &models.CategoryProductRequest{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		errors := helper.FormatValidationError(err)
		errBindJson := helper.NewUnprocessableEntityError(errors[0])
		response := helper.APIResponse(errBindJson.Message(), errBindJson.Status(), "error", nil)
		ctx.JSON(errBindJson.Status(), response)
		return
	}

	categoryProduct, err := cph.categoryProductService.UpdateCategoryProduct(ctx, id, input)
	if err != nil {
		response := helper.APIResponse(err.Message(), err.Status(), "error", nil)
		ctx.JSON(err.Status(), response)
		return
	}

	formattedResponse := helper.APIResponse("Category product updated successfully", http.StatusOK, "success", categoryProduct)
	ctx.JSON(http.StatusOK, formattedResponse)
}

func (cph *categoryProductHandler) DeleteCategoryProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	categoryProduct, err := cph.categoryProductService.DeleteCategoryProduct(ctx, id)
	if err != nil {
		response := helper.APIResponse(err.Message(), err.Status(), "error", nil)
		ctx.JSON(err.Status(), response)
		return
	}

	formattedResponse := helper.APIResponse("Category product deleted successfully", http.StatusOK, "success", categoryProduct)
	ctx.JSON(http.StatusOK, formattedResponse)
}
