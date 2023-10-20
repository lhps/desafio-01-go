package model

import (
	"github.com/asaskevich/govalidator"
	"time"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductRepositoryInterface interface {
	Create(product *Product) error
	FindAll() ([]*Product, error)
}

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey; autoIncrement" valid:"notnull"`
	Name        string    `json:"name" gorm:"type:varchar(255); not null" valid:"notnull"`
	Description string    `json:"description" gorm:"type:varchar(255); not null" valid:"notnull"`
	Price       float64   `json:"price" gorm:"type:float; not null" valid:"notnull"`
	CreatedAt   time.Time `json:"created_at" valid:"-"`
	UpdatedAt   time.Time `json:"updated_at" valid:"-"`
}

func (p *Product) isValid() error {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}

	return nil
}

func NewProduct(name string, description string, price float64) (*Product, error) {
	product := Product{
		Name:        name,
		Description: description,
		Price:       price,
	}

	product.CreatedAt = time.Now()

	err := product.isValid()
	if err != nil {
		return nil, err
	}

	return &product, nil
}
