package api

import (
	"github.com/gorilla/mux"
	"severyanochka/internals/app/controller/public"
	"severyanochka/internals/app/controller/public/help"
)

const (
	restApiPath = "/api/v1/"
	helpApiPath = "/help/"
)

// CreateRoute Пути, по которым доступен севрер
func CreateRoute(
	productController *public.ProductController,
	helpController *help.HelpController,
) *mux.Router {
	route := mux.NewRouter()
	//REST API
	route.HandleFunc(restApiPath+"page/preview-main", productController.PreviewMainPage).Methods("GET")
	route.HandleFunc(restApiPath+"products/search", productController.SearchByName).Methods("GET").Queries("q", "{q}")
	//HTML DOC API
	route.HandleFunc(helpApiPath+"main-page", helpController.ShowMainPageDoc)
	return route
}
