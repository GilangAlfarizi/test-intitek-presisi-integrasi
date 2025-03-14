package handler

import (
	"fullstack-test/helper"
	"fullstack-test/product"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService product.ServiceProduct
}

func NewProductHandler(productService product.ServiceProduct) *productHandler {
	return &productHandler{productService: productService}
}

func (h *productHandler) ReadProducts(ctx *gin.Context) {

	product, err := h.productService.GetProducts()

	if err != nil {
		err := helper.FailedResponse1(http.StatusNotFound, err.Error(), nil)
		ctx.AbortWithStatusJSON(err.Error.Code, err)
		return
	}

	response := helper.SuccessfulResponse1(product)
	ctx.JSON(http.StatusOK, response)
}

func (h *productHandler) ReadOneProduct(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errInvalidID := helper.FailedResponse1(http.StatusBadRequest, "Invalid product ID", idStr)
		ctx.AbortWithStatusJSON(errInvalidID.Error.Code, errInvalidID)
		return
	}

	product, err := h.productService.GetOneProduct(id)

	if err != nil {
		err := helper.FailedResponse1(http.StatusNotFound, err.Error(), nil)
		ctx.AbortWithStatusJSON(err.Error.Code, err)
		return
	}

	response := helper.SuccessfulResponse1(product)
	ctx.JSON(http.StatusOK, response)
}

func (h *productHandler) CreateProduct(ctx *gin.Context) {
	body := product.CreateProductDTO{}

	if err := ctx.Bind(&body); err != nil {
		errBindJson := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), body)
		ctx.AbortWithStatusJSON(errBindJson.Error.Code, errBindJson)
		return
	}

	// VALIDATION INPUT VALUE
	if err := body.Validate(); err != nil {
		errValidate := helper.FailedResponse1(http.StatusBadRequest, err.Error(), body)
		ctx.AbortWithStatusJSON(errValidate.Error.Code, errValidate)
		return
	}

	// productInput := product.ProductInput{Name: body.Name}
	productInput := product.ProductInput{
		Name:     body.Name,
		Sku:      body.Sku,
		Quantity: body.Quantity,
		Location: body.Location,
		Status:   body.Status,
	}
	product, errSvc := h.productService.CreateProduct(productInput)

	if errSvc != nil {
		errCreate := helper.FailedResponse1(http.StatusInternalServerError, errSvc.Error(), body)
		ctx.AbortWithStatusJSON(errCreate.Error.Code, errCreate)
		return
	}

	response := helper.SuccessfulResponse1(product)
	ctx.JSON(http.StatusOK, response)
}

func (h *productHandler) UpdateProductById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errInvalidID := helper.FailedResponse1(http.StatusBadRequest, "Invalid product ID", idStr)
		ctx.AbortWithStatusJSON(errInvalidID.Error.Code, errInvalidID)
		return
	}
	body := product.CreateProductDTO{}

	if err := ctx.Bind(&body); err != nil {
		errBindJson := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), body)
		ctx.AbortWithStatusJSON(errBindJson.Error.Code, errBindJson)
		return
	}

	// VALIDATION INPUT VALUE
	if err := body.Validate(); err != nil {
		errValidate := helper.FailedResponse1(http.StatusBadRequest, err.Error(), body)
		ctx.AbortWithStatusJSON(errValidate.Error.Code, errValidate)
		return
	}

	productInput := product.ProductInput{
		Name:     body.Name,
		Sku:      body.Sku,
		Quantity: body.Quantity,
		Location: body.Location,
		Status:   body.Status,
	}

	product, errSvc := h.productService.UpdateProduct(productInput, id)

	if errSvc != nil {
		errCreate := helper.FailedResponse1(http.StatusInternalServerError, errSvc.Error(), body)
		ctx.AbortWithStatusJSON(errCreate.Error.Code, errCreate)
		return
	}

	response := helper.SuccessfulResponse1(product)
	ctx.JSON(http.StatusOK, response)
}

func (h *productHandler) DeleteProductById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errInvalidID := helper.FailedResponse1(http.StatusBadRequest, "Invalid product ID", idStr)
		ctx.AbortWithStatusJSON(errInvalidID.Error.Code, errInvalidID)
		return
	}

	if _, err := h.productService.DeleteProduct(id); err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), id)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccessfulResponse1("product has successfully deleted")
	ctx.JSON(http.StatusOK, response)
}
