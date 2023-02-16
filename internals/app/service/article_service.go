package service

import (
	"severyanochka/internals/app/entity/site"
	"severyanochka/internals/app/repository"
)

// ArticleService Сервис для работы с site.Articles
type ArticleService struct {
	repository *repository.ArticleRepository
}

// NewArticleService Создает новый сервис для работы со статьями
func NewArticleService(repository *repository.ArticleRepository) *ArticleService {
	service := new(ArticleService)
	service.repository = repository
	return service
}

// GetAll Возвращает первые 20 статей, отсортированные по дате по убыванию
func (service *ArticleService) GetAll() []site.Articles {
	return service.repository.GetAll()
}

// FindById Получить статью по айди
func (service *ArticleService) FindById(id int64) site.Articles {
	return service.repository.FindById(id)
}
