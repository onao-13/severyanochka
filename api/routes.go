package api

import (
	"github.com/gorilla/mux"
	"severyanochka/internals/app/controller/public"
)

const (
	restApiPath = "/api/v1/"
)

// CreateRoute Пути, по которым доступен севрер
func CreateRoute(
	productController *public.ProductController,
	mainPageController *public.MainPageController,
	articleController *public.ArticleController,
	specialOfferController *public.SpecialOffersController,
	categoryController *public.CategoryController,
) *mux.Router {
	route := mux.NewRouter()
	//REST API
	route.HandleFunc(restApiPath+"page/home", mainPageController.PreviewMainPage).Methods("GET")
	route.HandleFunc(restApiPath+"products/search", productController.SearchByName).Methods("GET").Queries("q", "{q}")
	route.HandleFunc(restApiPath+"products/get/{id:[0-9]+}", productController.FindById).Methods("GET")
	route.HandleFunc(restApiPath+"articles/search/{id:[0-9]+}", articleController.FindById).Methods("GET")
	route.HandleFunc(restApiPath+"offers/search/{id:[0-9]+}", specialOfferController.FindById).Methods("GET")
	route.HandleFunc(restApiPath+"categories", categoryController.GetAll).Methods("GET")
	route.HandleFunc(restApiPath+"categories/{id:[0-9]+}/products", categoryController.FindCategoryProductsById).Methods("GET")
	return route
}
