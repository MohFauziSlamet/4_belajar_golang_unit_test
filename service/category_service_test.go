package service

import (
	"belajar_golang_unit_test/entity"
	"belajar_golang_unit_test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	// * PROGRAM MOCK NYA
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	//* test mock nya
	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	//* buat data successnya
	category := entity.Category{Id: "2", Name: "Handphone"}
	// * PROGRAM MOCK NYA
	categoryRepository.Mock.On("FindById", "2").Return(category)

	//* test mock nya
	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
