package repository

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"severyanochka/internals/app/entity/site"
)

// SpecialOffersRepository Репозиторий для работы с site.SpecialOffers
type SpecialOffersRepository struct {
	pool *pgxpool.Pool
}

// NewSpecialOffersRepository Создает новый репозиторий
func NewSpecialOffersRepository(pool *pgxpool.Pool) *SpecialOffersRepository {
	repository := new(SpecialOffersRepository)
	repository.pool = pool
	return repository
}

// GetAll Получение двух последних предложений
func (repository *SpecialOffersRepository) GetAll() []site.SpecialOffers {
	var offers []site.SpecialOffers

	query := "SELECT * FROM special_offers ORDER BY id DESC LIMIT 2"

	err := pgxscan.Select(context.Background(), repository.pool, &offers, query)
	if err != nil {
		log.Infoln("No special offer found. Error: ", err)
	}

	return offers
}

// FindById Поиск предложения по айди
func (repository *SpecialOffersRepository) FindById(id int64) site.SpecialOffers {
	var offer site.SpecialOffers

	query := "SELECT * FROM special_offers WHERE id = $1"

	err := pgxscan.Get(context.Background(), repository.pool, &offer, query, id)
	if err != nil {
		log.Infoln("No found special offer by id ", id, ". Error: ", err)
	}

	return offer
}
