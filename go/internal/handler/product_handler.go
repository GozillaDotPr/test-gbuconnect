package handler

import (
	"net/http"

	"github.com/app/gin-postgres-api/dto"
	"github.com/app/gin-postgres-api/internal/service"
	"github.com/app/gin-postgres-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type ProductHandlerImpl struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) ProductHandler {
	return &ProductHandlerImpl{productService: productService}
}

func (h *ProductHandlerImpl) GetAll(c *gin.Context) {
	products, err := h.productService.GetAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusOK, products)
}

func (h *ProductHandlerImpl) GetByID(c *gin.Context) {
	product, err := h.productService.GetByID(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}
	response.Success(c, http.StatusOK, product)
}

func (h *ProductHandlerImpl) Create(c *gin.Context) {
	var body dto.ProductRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		response.Error(c, http.StatusBadRequest, "name and price are required")
		return
	}

	userID, _ := c.Get("username")

	product, err := h.productService.Create(body.Name, body.Desc, body.Price, userID.(string))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, product)
}

func (h *ProductHandlerImpl) Update(c *gin.Context) {
	var body dto.ProductRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.productService.Update(c.Param("id"), body.Name, body.Desc, body.Price)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "product not found" {
			status = http.StatusNotFound
		}
		response.Error(c, status, err.Error())
		return
	}

	response.Success(c, http.StatusOK, product)
}

func (h *ProductHandlerImpl) Delete(c *gin.Context) {
	if err := h.productService.Delete(c.Param("id")); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "product not found" {
			status = http.StatusNotFound
		}
		response.Error(c, status, err.Error())
		return
	}
	response.Success(c, http.StatusOK, nil)
}
