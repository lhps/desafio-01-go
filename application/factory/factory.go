package factory

import (
	"github.com/jinzhu/gorm"
	"github.com/lhps/desafio-01/application/usecase"
	"github.com/lhps/desafio-01/infrastructure/repository"
)

func ProductUseCaseFactory(database *gorm.DB) usecase.ProductUseCase {
	productRepository := repository.ProductRepositoryDb{
		Db: database,
	}

	productUseCase := usecase.ProductUseCase{ProductRepository: &productRepository}

	return productUseCase
}
