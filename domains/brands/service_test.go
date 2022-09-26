package brands

import (
	"errors"
	"testing"
	"tirenn/catalyst/dto"
	"tirenn/catalyst/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Init() (*RepositoryMock, Service) {
	var repo = &RepositoryMock{Mock: mock.Mock{}}
	return repo, Service{repository: repo}
}

func TestCreateSuccess(t *testing.T) {
	repo, service := Init()

	createBrand := dto.CreateBrand{
		Name: "Brand A",
	}

	brand := models.Brand{
		ID:   1,
		Name: "Brand A",
	}

	repo.Mock.On("Create", mock.AnythingOfType("*models.Brand")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*models.Brand)
		*arg = brand
	})

	res, err := service.Create(createBrand)
	assert.Equal(t, nil, err)
	assert.Equal(t, createBrand.Name, res.Name)
}

func TestCreateFail(t *testing.T) {
	repo, service := Init()

	createBrand := dto.CreateBrand{
		Name: "Brand A",
	}

	errReturn := errors.New("create error")
	repo.Mock.On("Create", mock.AnythingOfType("*models.Brand")).Return(errReturn)

	res, err := service.Create(createBrand)
	assert.Equal(t, errReturn, err)
	assert.NotEqual(t, createBrand.Name, res.Name)
	assert.Empty(t, res)
}
