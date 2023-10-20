package usecase

import (
	"github.com/lhps/desafio-01/domain/model"
)

type ProductUseCase struct {
	ProductRepository model.ProductRepositoryInterface
}

func (p *ProductUseCase) CreateProduct(name string, description string, price float64) (*model.Product, error) {
	product, err := model.NewProduct(name, description, price)
	if err != nil {
		return nil, err
	}

	err = p.ProductRepository.Create(product)
	if product.ID == 0 {
		return nil, err
	}

	return product, nil
}

func (p *ProductUseCase) ListProducts() ([]*model.Product, error) {
	var products []*model.Product
	products, err := p.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}
