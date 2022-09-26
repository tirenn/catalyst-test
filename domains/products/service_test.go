package products

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"tirenn/catalyst/dto"
	"tirenn/catalyst/models"
)

func Init() (*RepositoryMock, Service) {
	var repo = &RepositoryMock{Mock: mock.Mock{}}
	return repo, Service{repository: repo}
}

func TestCreateSuccess(t *testing.T) {
	repo, service := Init()

	createProduct := dto.CreateProduct{
		Name:    "Product A",
		Price:   1000,
		BrandID: 1,
	}

	brand := models.Brand{
		ID:   1,
		Name: "Brand A",
	}

	product := models.Product{
		ID:    1,
		Name:  "Product A",
		Price: 1000,
		Brand: brand,
	}

	repo.Mock.On("Create", mock.AnythingOfType("*models.Product")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*models.Product)
		*arg = product
	})

	res, err := service.Create(createProduct)
	assert.Equal(t, nil, err)
	assert.Equal(t, createProduct.Name, res.Name)

}

func TestCreateFail(t *testing.T) {
	repo, service := Init()

	createProduct := dto.CreateProduct{
		Name:  "Product A",
		Price: 1000,
	}
	errReturn := errors.New("create error")
	repo.Mock.On("Create", mock.AnythingOfType("*models.Product")).Return(errReturn)

	res, err := service.Create(createProduct)
	assert.Equal(t, errReturn, err)
	assert.NotEqual(t, createProduct.Name, res.Name)
	assert.Empty(t, res)
}

func TestGetSuccess(t *testing.T) {
	repo, service := Init()

	var id int64
	id = 1

	brand := models.Brand{
		ID:   1,
		Name: "Brand A",
	}

	product := models.Product{
		ID:    1,
		Name:  "Product A",
		Price: 1000,
		Brand: brand,
	}

	repo.Mock.On("Get", id).Return(product, nil)

	res, err := service.Get(id)
	assert.Equal(t, nil, err)
	assert.Equal(t, id, res.ID)
	assert.Equal(t, brand.ID, res.Brand.ID)
}

func TestGetFail(t *testing.T) {
	repo, service := Init()

	var id int64
	id = 1

	product := models.Product{}

	errReturn := errors.New("get error")
	repo.Mock.On("Get", id).Return(product, errReturn)

	res, err := service.Get(id)
	assert.Equal(t, errReturn, err)
	assert.NotEqual(t, id, res.ID)
}

func TestGetByBrandSuccess(t *testing.T) {
	repo, service := Init()

	var brandID int64
	brandID = 1

	brand := models.Brand{
		ID:   1,
		Name: "Brand A",
	}

	products := []models.Product{
		{
			ID:    1,
			Name:  "Product A",
			Price: 1000,
			Brand: brand,
		},
		{
			ID:    2,
			Name:  "Product B",
			Price: 2000,
			Brand: brand,
		},
	}

	repo.Mock.On("GetByBrand", brandID).Return(products, nil)

	res, err := service.GetByBrand(brandID)
	assert.Equal(t, nil, err)
	assert.Equal(t, products[0].ID, res[0].ID)
	assert.Equal(t, len(products), len(res))

}

func TestGetByBrandFail(t *testing.T) {
	repo, service := Init()

	var brandID int64
	brandID = 1

	var products []models.Product

	errReturn := errors.New("get error")
	repo.Mock.On("GetByBrand", brandID).Return(products, errReturn)

	res, err := service.GetByBrand(brandID)
	assert.Equal(t, errReturn, err)
	assert.Equal(t, 0, len(res))
}
