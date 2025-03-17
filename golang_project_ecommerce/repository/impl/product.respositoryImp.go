// Package implRepository...
package implRepository

import (
	"golang_project_ecommerce/models"
	"golang_project_ecommerce/repository"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

// NewProductRepositoryImpl is
func NewProductRepositoryImpl(Db *gorm.DB) repository.ProductRepository {
	return &ProductRepositoryImpl{Db: Db}
}

// Create repository imp
func (p ProductRepositoryImpl) Save(product models.Product) {

}

// Update repository imp
func (p ProductRepositoryImpl) Update(product models.Product) {

}

// Delete repository imp
func (p ProductRepositoryImpl) Delete(productId int) {

}

// FindById repository imp
func (p ProductRepositoryImpl) FindById(productId int) (models.Product, error) {
	return models.Product{}, nil

}

// FindAll repository imp
func (p ProductRepositoryImpl) FindAll() []models.Product {
	return []models.Product{}

}
