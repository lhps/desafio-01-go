package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/lhps/desafio-01/domain/model"
)

type ProductRepositoryDb struct {
	Db *gorm.DB
}

func (p *ProductRepositoryDb) Create(product *model.Product) error {
	err := p.Db.Create(product).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepositoryDb) FindAll() ([]*model.Product, error) {
	var products []*model.Product

	p.Db.Find(&products)

	return products, nil
}
