package controller

import (
	"go-restapi/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// doc: primeira letra minúscula define que a estrutura ou função é visível somente dentro do pacote
type productController struct {
	// usecase
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}