package public

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"severyanochka/internals/app/handler/rest"
	"severyanochka/internals/app/service"
	"strconv"
)

type ArticleController struct {
	service *service.ArticleService
}

var log = logrus.New()

func NewArticleController(service *service.ArticleService) *ArticleController {
	controller := new(ArticleController)
	controller.service = service
	return controller
}

func (controller *ArticleController) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars["id"] == "" {
		rest.WrapError(errors.New("Поле ID пустое"), w)
		return
	}

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Errorln("Error parse vars. Error: ", err)
		return
	}

	article := controller.service.FindById(id)

	if article.Id == 0 {
		rest.WrapNotFound(w)
		return
	}

	rest.WrapOk(article, w)
}
