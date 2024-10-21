package controller

import (
	"go-api-commerce/model"
	"go-api-commerce/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID:    1,
			Name:  "Teclado Lakshmi",
			Price: 200,
		},
	}

	ctx.JSON(http.StatusOK, products)
}
