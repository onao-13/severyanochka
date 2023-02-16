package service

import (
	"severyanochka/internals/app/entity/site"
	"severyanochka/internals/app/repository"
)

// SpecialOffersService Сервис для работы с site.SpecialOffers
type SpecialOffersService struct {
	repository *repository.SpecialOffersRepository
}

// NewSpecialOffersService Создает новый сервис
func NewSpecialOffersService(repository *repository.SpecialOffersRepository) *SpecialOffersService {
	service := new(SpecialOffersService)
	service.repository = repository
	return service
}

// GetAll Возврщает последние 2 предложения
func (service *SpecialOffersService) GetAll() []site.SpecialOffers {
	return service.repository.GetAll()
}

// FindById Поиск предложения по айди
func (service *SpecialOffersService) FindById(id int64) site.SpecialOffers {
	return service.repository.FindById(id)
}
