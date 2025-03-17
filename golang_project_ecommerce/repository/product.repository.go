// Package repository ...
package repository

import "golang_project_ecommerce/models"

// ProductRepository is
type ProductRepository interface {
	Save(products models.Product)
	Update(tags models.Product)
	Delete(tagsId int)
	FindById(tagsId int) (tags models.Product, err error)
	FindAll() []models.Product
}
