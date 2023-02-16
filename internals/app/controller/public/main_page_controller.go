package public

import (
	"net/http"
	"severyanochka/internals/app/handler/rest"
	"severyanochka/internals/app/service"
)

// MainPageController Контроллер для вывода всей информации главной страницы
type MainPageController struct {
	articleService       *service.ArticleService
	productService       *service.ProductService
	specialOffersService *service.SpecialOffersService
}

// NewMainPageController Создает новый контроллер главной страницы
func NewMainPageController(
	articleService *service.ArticleService,
	productService *service.ProductService,
	specialOffersService *service.SpecialOffersService,
) *MainPageController {
	controller := new(MainPageController)
	controller.articleService = articleService
	controller.productService = productService
	controller.specialOffersService = specialOffersService
	return controller
}

//TODO: Добавить подсчет процента по карте магазина

// PreviewMainPage Контроллер для отправки всех данных главной страницы.
// Содержит:
//
// - список акций
//
// - список новинок
//
// - список последних покупок пользователя
//
// - специальные предложения
//
// - статьи
func (controller *MainPageController) PreviewMainPage(w http.ResponseWriter, r *http.Request) {
	var result = map[string]interface{}{
		"stock":         controller.productService.GetStock(),
		"newProducts":   controller.productService.GetNewProducts(),
		"articles":      controller.articleService.GetAll(),
		"specialOffers": controller.specialOffersService.GetAll(),
	}
	rest.WrapOk(result, w)
}
