package brands

import (
	"github.com/stretchr/testify/mock"
	"tirenn/catalyst/models"
)

type RepositoryMock struct {
	Mock mock.Mock
}

func (r *RepositoryMock) Create(brand *models.Brand) (err error) {
	args := r.Mock.Called(brand)
	return args.Error(0)
}
