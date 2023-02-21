package public

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"severyanochka/internals/app/handler/rest"
	"severyanochka/internals/app/service"
	"severyanochka/internals/utils/parser"
)

// ArticleController Контроллер для работы со статьями
type ArticleController struct {
	service *service.ArticleService
	parser  *parser.RequestParserVars
}

var log = logrus.New()

// NewArticleController Создает новый контроллер статей
func NewArticleController(service *service.ArticleService, parser *parser.RequestParserVars) *ArticleController {
	controller := new(ArticleController)
	controller.service = service
	controller.parser = parser
	return controller
}

// FindById Поиск статьи по айди
func (controller *ArticleController) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars["id"] == "" {
		rest.WrapError(errors.New("Поле ID пустое"), w)
		return
	}

	id := controller.parser.ParseToInt64(r, "id")

	article := controller.service.FindById(id)

	if article.Id == 0 {
		rest.WrapNotFound(w)
		return
	}

	rest.WrapOk(article, w)
}
