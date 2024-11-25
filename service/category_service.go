package service

import (
	"belajar_golang_unit_test/entity"
	"belajar_golang_unit_test/repository"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)

	if category != nil {
		return category, nil
	}

	return nil, errors.New("category not found")
}
