package public

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"severyanochka/internals/app/handler/rest"
	"severyanochka/internals/app/service"
	"severyanochka/internals/utils/parser"
)

type CategoryController struct {
	service *service.CategoryService
	parser  *parser.RequestParserVars
}

func NewCategoryController(service *service.CategoryService, parser *parser.RequestParserVars) *CategoryController {
	controller := new(CategoryController)
	controller.service = service
	controller.parser = parser
	return controller
}

func (controller *CategoryController) GetAll(w http.ResponseWriter, r *http.Request) {
	categories := controller.service.GetAll()

	if categories == nil {
		rest.WrapNotFound(w)
		return
	}

	rest.WrapOk(categories, w)
}

func (controller *CategoryController) FindCategoryProductsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars["id"] == "" {
		rest.WrapError(errors.New("Поле ID пустое"), w)
		return
	}

	id := controller.parser.ParseToInt64(r, "id")

	result := controller.service.GetCategoryProductsById(id)

	if result["products"] == "" || result["categoryName"] == "" {
		rest.WrapNotFound(w)
		return
	}

	rest.WrapOk(result, w)
}
