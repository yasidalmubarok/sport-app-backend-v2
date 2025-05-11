package handlers

import (
	"net/http"
	"sport-app-backend/helper"
	"sport-app-backend/models"
	"sport-app-backend/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) *productHandler {
	return &productHandler{productService: productService}
}

func (ph *productHandler) CreateProduct(c *gin.Context) {
	// Ambil file gambar
	file, err := c.FormFile("image")
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Proses upload dan kompresi gambar
	imagePath, err := helper.UploadAndCompressImage(file, 300)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Bind data form
	priceSell, _ := strconv.ParseFloat(c.PostForm("price_sell"), 64)
	priceBuy, _ := strconv.ParseFloat(c.PostForm("price_buy"), 64)
	stock, _ := strconv.Atoi(c.PostForm("stock"))

	req := models.CreateProductRequest{
		Name:      c.PostForm("name"),
		Category:  c.PostForm("category"),
		PriceSell: priceSell,
		PriceBuy:  priceBuy,
		Stock:     stock,
		Status:    c.PostForm("status"),
		Image:     imagePath,
	}

	// Panggil service
	product, err := ph.productService.CreateProduct(c.Request.Context(), &req)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Product created successfully", http.StatusCreated, "success", product)
	c.JSON(http.StatusCreated, response)
}

func (ph *productHandler) GetProducts(ctx *gin.Context) {
	products, err := ph.productService.GetAllProduct(ctx)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Products found successfully", http.StatusOK, "success", products)
	ctx.JSON(http.StatusOK, response)
}

func (ph *productHandler) GetProductByID(ctx *gin.Context) {
	id := ctx.Param("id")

	product, err := ph.productService.GetProductByID(ctx, id)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Product found successfully", http.StatusOK, "success", product)
	ctx.JSON(http.StatusOK, response)
}

func (ph *productHandler) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	// Ambil file gambar
	file, err := ctx.FormFile("image")
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Proses upload dan kompresi gambar
	imagePath, err := helper.UploadAndCompressImage(file, 300)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Bind data form
	priceSell, _ := strconv.ParseFloat(ctx.PostForm("price_sell"), 64)
	priceBuy, _ := strconv.ParseFloat(ctx.PostForm("price_buy"), 64)
	stock, _ := strconv.Atoi(ctx.PostForm("stock"))

	req := &models.CreateProductRequest{
		Name:      ctx.PostForm("name"),
		Category:  ctx.PostForm("category"),
		PriceSell: priceSell,
		PriceBuy:  priceBuy,
		Stock:     stock,
		Status:    ctx.PostForm("status"),
		Image:     imagePath,
	}

	// Panggil service
	product, err := ph.productService.UpdateProduct(ctx, id, req)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Product updated successfully", http.StatusOK, "success", product)
	ctx.JSON(http.StatusOK, response)
}

func (ph *productHandler) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	product, err := ph.productService.DeleteProduct(ctx, id)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Product deleted successfully", http.StatusOK, "success", product)
	ctx.JSON(http.StatusOK, response)
}
