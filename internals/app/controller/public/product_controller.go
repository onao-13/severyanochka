package public

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"severyanochka/internals/app/entity/net/response"
	"severyanochka/internals/app/handler/rest"
	"severyanochka/internals/app/service"
	"severyanochka/internals/utils/parser"
)

// ProductController Контроллер для работы с продуктами
type ProductController struct {
	service *service.ProductService
	parser  *parser.RequestParserVars
}

// NewProductController Создание контроллера продуктов
func NewProductController(service *service.ProductService, parser *parser.RequestParserVars) *ProductController {
	controller := new(ProductController)
	controller.service = service
	controller.parser = parser
	return controller
}

// SearchByName Поиск продукта по названию
func (controller *ProductController) SearchByName(w http.ResponseWriter, r *http.Request) {
	var products []response.ProductResponse

	name := r.URL.Query().Get("q")

	if name == "" {
		rest.WrapError(errors.New("Параметр q пустой. Введите запрос для поиска"), w)
		return
	}

	products = controller.service.SearchProductByName(name)

	if len(products) == 0 {
		rest.WrapNotFound(w)
	} else {
		rest.WrapOk(products, w)
	}
}

func (controller *ProductController) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars["id"] == "" {
		rest.WrapError(errors.New("Поле ID пустое"), w)
		return
	}

	id := controller.parser.ParseToInt64(r, "id")

	product := controller.service.FindById(id)

	if product.Id <= 0 {
		rest.WrapNotFound(w)
		return
	}

	rest.WrapOk(product, w)
}
