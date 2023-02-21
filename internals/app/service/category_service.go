package service

import (
	"severyanochka/internals/app/entity/product"
	"severyanochka/internals/app/repository"
)

type CategoryService struct {
	categoryRepository *repository.CategoryRepository
	productService     *ProductService
}

func NewCategoryService(categoryRepository *repository.CategoryRepository, productService *ProductService) *CategoryService {
	service := new(CategoryService)
	service.categoryRepository = categoryRepository
	service.productService = productService
	return service
}

func (service *CategoryService) GetAll() []product.Categories {
	return service.categoryRepository.GetAll()
}

func (service *CategoryService) GetCategoryProductsById(id int64) map[string]interface{} {
	var categoryProducts = map[string]interface{}{
		"categoryName": service.categoryRepository.FindCategoryNameById(id),
		"products":     service.productService.FindAllProductsByCategoryId(id),
	}
	return categoryProducts
}
