package server

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"net/http"
	"severyanochka/api"
	"severyanochka/internals/app/controller/public"
	help "severyanochka/internals/app/controller/public/help"
	"severyanochka/internals/app/repository"
	"severyanochka/internals/app/service"
	"severyanochka/internals/config"
	"severyanochka/internals/utils/db"
	"severyanochka/internals/utils/md"
	"time"
)

// Server Сервер с конфигурацией и пулом для подключения к бд
type Server struct {
	cfg     config.Config
	ctx     context.Context
	srv     *http.Server
	pool    *pgxpool.Pool
	dbUtils *db.TableUtils
}

var log = logrus.New()

// NewServer Создание нового сервера
func NewServer(ctx context.Context, cfg config.Config) *Server {
	server := new(Server)
	server.ctx = ctx
	server.cfg = cfg
	return server
}

// Start Запуск сервера
func (server *Server) Start() {
	log.Infoln("Server starting")

	var err error

	server.pool, err = pgxpool.Connect(server.ctx, server.cfg.DbUrl())
	if err != nil {
		log.Fatalln("Error connecting to database: ", err)
	}

	server.dbUtils = db.New(server.pool, server.ctx)
	initializeDb(server.dbUtils)

	defer server.pool.Close()

	productRepository := repository.NewProductRepository(server.pool)
	articleRepository := repository.NewArticleRepository(server.pool)
	specialOffersRepository := repository.NewSpecialOffersRepository(server.pool)

	productService := service.NewProductService(productRepository)
	articleService := service.NewArticleService(articleRepository)
	specialOffersService := service.NewSpecialOffersService(specialOffersRepository)

	productController := public.NewProductController(productService)
	helpController := help.New(md.New())
	mainPageController := public.NewMainPageController(articleService, productService, specialOffersService)
	articleController := public.NewArticleController(articleService)
	specialOffersController := public.NewSpecialOffersController(specialOffersService)

	routes := api.CreateRoute(productController, helpController, mainPageController, articleController, specialOffersController)

	server.srv = &http.Server{
		Addr:    ":" + server.cfg.Port,
		Handler: routes,
	}

	log.Infoln("Server started")

	err = server.srv.ListenAndServe()
	if err != nil {
		log.Fatalln("Error listen server: ", err)
	}
}

// Shutdown Остановка сервера
func (server *Server) Shutdown() {
	log.Infoln("Server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	dropTables(server.dbUtils)

	defer func() {
		cancel()
	}()

	var err error

	if err = server.srv.Shutdown(ctx); err != nil {
		log.Fatalln("Server shutdown failed: ", err)
	}

	if err == http.ErrServerClosed {
		err = nil
	}
}

func initializeDb(utils *db.TableUtils) {
	utils.CreateTables()
	utils.UploadDevData()
}

func dropTables(utils *db.TableUtils) {
	utils.DropTables()
}
