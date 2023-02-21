package repository

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"severyanochka/internals/app/entity/product"
)

type CategoryRepository struct {
	pool *pgxpool.Pool
}

func NewCategoryRepository(pool *pgxpool.Pool) *CategoryRepository {
	reposirory := new(CategoryRepository)
	reposirory.pool = pool
	return reposirory
}

func (repository *CategoryRepository) GetAll() []product.Categories {
	var categories []product.Categories

	query := "SELECT * FROM categories"

	err := pgxscan.Select(context.Background(), repository.pool, &categories, query)
	if err != nil {
		log.Infoln("No categories found. Error: ", err)
	}

	return categories
}

func (repository *CategoryRepository) FindCategoryNameById(id int64) string {
	var category string

	query := "SELECT name FROM categories WHERE id = $1"

	err := pgxscan.Get(context.Background(), repository.pool, &category, query, id)
	if err != nil {
		log.Infoln("No category found with ", id, " id. Error: ", err)
	}

	return category
}
