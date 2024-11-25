package repository

import "belajar_golang_unit_test/entity"

type CategoryRepository interface {
	FindById(Id string) *entity.Category
}
