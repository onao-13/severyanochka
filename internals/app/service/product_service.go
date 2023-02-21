package service

import (
	"github.com/sirupsen/logrus"
	"severyanochka/internals/app/entity/net/response"
	"severyanochka/internals/app/entity/product"
	"severyanochka/internals/app/repository"
)

// ProductService Сервис для работы с продуктами
type ProductService struct {
	repository *repository.ProductRepository
}

var log = logrus.New()

// NewProductService Создание сервиса продуктов
func NewProductService(repository *repository.ProductRepository) *ProductService {
	service := new(ProductService)
	service.repository = repository
	return service
}

// TODO: FIX THIS (send null)
// GetStock Сервис для отображения акционных продуктов на главной /*
func (service *ProductService) GetStock() []response.ProductResponse {
	var products []response.ProductResponse

	for _, p := range service.repository.FindStock() {
		products = append(products, createProductResponse(p))
	}

	return products
}

// GetNewProducts Сервис на получение новых продуктов на главной/*
func (service *ProductService) GetNewProducts() []response.ProductResponse {
	var products []response.ProductResponse

	for _, p := range service.repository.FindNewProduct() {
		products = append(products, createProductResponse(p))
	}

	return products
}

// GetBoughtBefore Сервис для получение покупок,которые раньше совершал пользователь
// на главной
func (service *ProductService) GetBoughtBefore() []response.ProductResponse {
	// TODO: UPDATE THIS
	return nil
}

// SearchProductByName Сервис для поиска продукта по названию /*
func (service *ProductService) SearchProductByName(name string) []response.ProductResponse {
	var products []response.ProductResponse

	result, err := service.repository.SearchByName(name)
	if err != nil {
		return products
	}

	for _, p := range result {
		products = append(products, createProductResponse(p))
	}

	return products
}

func (service *ProductService) FindAllProductsByCategoryId(id int64) []response.ProductResponse {
	var products []response.ProductResponse

	for _, p := range service.repository.FindAllProductsByCategoryId(id) {
		products = append(products, createProductResponse(p))
	}

	return products
}

func (service *ProductService) FindById(id int64) product.Products {
	return service.repository.FindById(id)
}

// Конвертирование product.Products в response.ProductResponse
func createProductResponse(product product.Products) response.ProductResponse {
	var result response.ProductResponse
	result.Id = product.Id
	result.Name = product.Name.String
	result.Price = product.Price
	result.ImageUrl = product.ImageUrl.String
	result.County = product.Country.String
	result.Sale = product.Sale
	result.Rating = product.Rating

	return result
}
