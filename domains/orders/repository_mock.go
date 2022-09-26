package orders

import (
	"github.com/stretchr/testify/mock"
	"tirenn/catalyst/models"
)

type RepositoryMock struct {
	Mock mock.Mock
}

func (r *RepositoryMock) Create(order *models.Order) (err error) {
	args := r.Mock.Called(order)
	return args.Error(0)
}

func (r *RepositoryMock) GetProduct(id int64) (product models.Product, err error) {
	args := r.Mock.Called(id)
	result := args.Get(0)
	return result.(models.Product), args.Error(1)
}

func (r *RepositoryMock) Get(id int64) (order models.Order, err error) {
	args := r.Mock.Called(id)
	result := args.Get(0)
	return result.(models.Order), args.Error(1)
}
