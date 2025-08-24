package usecase

import (
	"go-restapi/model"
	"go-restapi/repository"
)

// doc: primeira letra maiúscula define que a estrutura ou função é visível fora do pacote
type ProductUsecase struct {
	// repository
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

// função que trata regras de negócio para rota get /products
func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}