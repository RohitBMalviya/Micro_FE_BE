// Package implService ...
package implService

import (
	"golang_project_ecommerce/repository"
	"golang_project_ecommerce/services"

	"github.com/go-playground/validator/v10"
)

// ProductServiceImpl is
type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	Validate          *validator.Validate
}

// NewProductServiceImpl is
func NewProductServiceImpl(tagRepository repository.ProductRepository, validate *validator.Validate) services.ProductService {
	return &ProductServiceImpl{
		ProductRepository: tagRepository,
		Validate:          validate,
	}
}

// Create service impl
func (p ProductServiceImpl) Create(product string) {

}

// Update service impl
func (p ProductServiceImpl) Update(product string) {

}

// Delete service impl
func (p ProductServiceImpl) Delete(productId int) {

}

// FindById service impl
func (p ProductServiceImpl) FindById(productId int) string {
	return ""

}

// FindAll service impl
func (p ProductServiceImpl) FindAll() []string {
	return []string{""}

}
