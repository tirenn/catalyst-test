package products

import (
	"github.com/stretchr/testify/mock"
	"tirenn/catalyst/models"
)

type RepositoryMock struct {
	Mock mock.Mock
}

func (r *RepositoryMock) Create(product *models.Product) (err error) {
	args := r.Mock.Called(product)
	return args.Error(0)
}

func (r *RepositoryMock) Get(id int64) (product models.Product, err error) {
	args := r.Mock.Called(id)
	result := args.Get(0)
	return result.(models.Product), args.Error(1)
}

func (r *RepositoryMock) GetByBrand(brandID int64) (products []models.Product, err error) {
	args := r.Mock.Called(brandID)
	result := args.Get(0)
	return result.([]models.Product), args.Error(1)
}
