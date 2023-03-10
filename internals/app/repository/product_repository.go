package repository

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"severyanochka/internals/app/entity/product"
)

// ProductRepository Репозиторий для работы с базой данных продуктов
type ProductRepository struct {
	pool *pgxpool.Pool
}

// NewProductRepository Создание репозитория продуктов
func NewProductRepository(pool *pgxpool.Pool) *ProductRepository {
	repository := new(ProductRepository)
	repository.pool = pool
	return repository
}

// FindStock Получение списка акционных продуктов в размере 20 штук
func (repository *ProductRepository) FindStock() []product.Products {
	var products []product.Products

	//TODO: ВЫНЕСТИ ЗАПРОСЫ В ФАЙЛ SQL
	query := "SELECT id, name, image_url, price, sale, rating, country FROM products WHERE sale > 0 LIMIT 20"

	err := pgxscan.Select(context.Background(), repository.pool, &products, query)
	if err != nil {
		log.Infoln("No stock is db. Error: ", err)
	}

	return products
}

// FindNewProduct Получение последних добавленных продуктов в базу в размере 20 штук
func (repository *ProductRepository) FindNewProduct() []product.Products {
	var products []product.Products

	query := "SELECT * FROM products ORDER BY id DESC LIMIT 20"

	err := pgxscan.Select(context.Background(), repository.pool, &products, query)
	if err != nil {
		log.Infoln("No products in db. Error: ", err)
	}

	return products
}

func FindBoughtBefore() {
	//TODO: ADD TABLE FOR SAVE USER BUYING PRODUCTS
}

func (repository *ProductRepository) SearchByName(name string) ([]product.Products, error) {
	var products []product.Products

	query := "SELECT id, name, image_url, price, sale, rating, country FROM products WHERE name LIKE $1;"

	err := pgxscan.Select(context.Background(), repository.pool, &products, query, name+"%")
	if err != nil {
		log.Infoln("Product by name ", name, " not found, Error: ", err)
	}

	return products, err
}

func (repository *ProductRepository) FindById(id int64) product.Products {
	var prod product.Products

	query := "SELECT * FROM products WHERE id = $1"

	err := pgxscan.Get(context.Background(), repository.pool, &prod, query, id)
	if err != nil {
		log.Infoln("No products found. Error: ", err)
	}

	return prod
}

func (repository *ProductRepository) FindAllProductsByCategoryId(id int64) []product.Products {
	var products []product.Products

	query := "SELECT id, name, image_url, price, sale, rating, country FROM products WHERE category_id = $1"

	err := pgxscan.Select(context.Background(), repository.pool, &products, query, id)
	if err != nil {
		log.Infoln("No found products by category id ", id, ". Error: ", err)
	}

	return products
}
