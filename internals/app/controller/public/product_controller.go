package public

import (
	"net/http"
	"severyanochka/internals/app/entity/net/response"
	"severyanochka/internals/app/handler/rest"
	"severyanochka/internals/app/service"
)

// ProductController Контроллер для работы с продуктами
type ProductController struct {
	service *service.ProductService
}

// NewProductController Создание контроллера продуктов
func NewProductController(service *service.ProductService) *ProductController {
	controller := new(ProductController)
	controller.service = service
	return controller
}

// SearchByName Поиск продукта по названию
func (controller *ProductController) SearchByName(w http.ResponseWriter, r *http.Request) {
	var products []response.ProductResponse

	name := r.URL.Query().Get("q")
	products = controller.service.SearchProductByName(name)

	if len(products) == 0 {
		rest.WrapNotFound(w)
	} else {
		rest.WrapOk(products, w)
	}
}
