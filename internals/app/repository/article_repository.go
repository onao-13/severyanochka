package repository

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"severyanochka/internals/app/entity/site"
)

// ArticleRepository Репозиторий для работы в бд со site.Articles
type ArticleRepository struct {
	pool *pgxpool.Pool
}

var log = logrus.New()

// NewArticleRepository Создает новый репозиторий статей
func NewArticleRepository(pool *pgxpool.Pool) *ArticleRepository {
	repository := new(ArticleRepository)
	repository.pool = pool
	return repository
}

// GetAll Получить первые 20 статей, отсортированные по дате по убыванию
func (repository *ArticleRepository) GetAll() []site.Articles {
	var articles []site.Articles

	query := "SELECT * FROM articles ORDER BY date DESC LIMIT 20"

	err := pgxscan.Select(context.Background(), repository.pool, &articles, query)
	if err != nil {
		log.Infoln("No articles found. Error: ", err)
	}

	return articles
}

// FindById Поиск статьи по айди
func (repository *ArticleRepository) FindById(id int64) site.Articles {
	var article site.Articles

	query := "SELECT * FROM articles WHERE id = $1"

	err := pgxscan.Get(context.Background(), repository.pool, &article, query, id)
	if err != nil {
		log.Infoln("No articles found by id ", id, ". Error: ", err)
	}

	return article
}
