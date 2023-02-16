package public

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"severyanochka/internals/app/handler/rest"
	"severyanochka/internals/app/service"
	"strconv"
)

type SpecialOffersController struct {
	service *service.SpecialOffersService
}

func NewSpecialOffersController(service *service.SpecialOffersService) *SpecialOffersController {
	controller := new(SpecialOffersController)
	controller.service = service
	return controller
}

func (controller *SpecialOffersController) FindById(w http.ResponseWriter, r *http.Request) {
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

	offer := controller.service.FindById(id)

	if offer.Id == 0 {
		rest.WrapNotFound(w)
		return
	}

	rest.WrapOk(offer, w)
}
