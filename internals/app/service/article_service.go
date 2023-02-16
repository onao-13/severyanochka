package service

import (
	"severyanochka/internals/app/entity/site"
	"severyanochka/internals/app/repository"
)

type ArticleService struct {
	repository *repository.ArticleRepository
}

func NewArticleService(repository *repository.ArticleRepository) *ArticleService {
	service := new(ArticleService)
	service.repository = repository
	return service
}

func (service *ArticleService) GetAll() []site.Articles {
	return service.repository.GetAll()
}

func (service *ArticleService) FindById(id int64) site.Articles {
	return service.repository.FindById(id)
}
