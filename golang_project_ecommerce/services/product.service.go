// Package services ...
package services

// ProductService interface
type ProductService interface {
	Create(tags string)
	Update(tags string)
	Delete(tagsId int)
	FindById(tagsId int) string
	FindAll() []string
}
