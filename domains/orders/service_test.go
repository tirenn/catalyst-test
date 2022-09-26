package orders

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

	createOrderProducts := []dto.CreateOrderProduct{
		{
			ProductID: 1,
			Amount:    6,
		},
	}

	createOrder := dto.CreateOrder{
		OrderProducts: createOrderProducts,
	}

	product := models.Product{
		ID:    1,
		Name:  "Product A",
		Price: 1000,
	}

	orderProducts := []models.OrderProduct{
		{
			Product:      product,
			ProductPrice: product.Price,
			Amount:       6,
			Total:        6000,
		},
	}

	order := models.Order{
		ID:            1,
		Total:         6000,
		OrderProducts: orderProducts,
	}

	repo.Mock.On("GetProduct", createOrderProducts[0].ProductID).Return(product, nil)
	repo.Mock.On("Create", mock.AnythingOfType("*models.Order")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*models.Order)
		*arg = order
	})

	res, err := service.Create(createOrder.OrderProducts)
	assert.Equal(t, nil, err)
	assert.Equal(t, createOrder.OrderProducts[0].ProductID, res.OrderProducts[0].Product.ID)

}

func TestCreateFail(t *testing.T) {
	repo, service := Init()

	createOrderProducts := []dto.CreateOrderProduct{
		{
			ProductID: 1,
			Amount:    6,
		},
	}

	createOrder := dto.CreateOrder{
		OrderProducts: createOrderProducts,
	}

	product := models.Product{}
	repo.Mock.On("GetProduct", createOrderProducts[0].ProductID).Return(product, nil)

	errReturn := errors.New("create error")
	repo.Mock.On("Create", mock.AnythingOfType("*models.Order")).Return(errReturn)

	res, err := service.Create(createOrder.OrderProducts)
	assert.Equal(t, errReturn, err)
	assert.Empty(t, res)
}

func TestGetSuccess(t *testing.T) {
	repo, service := Init()

	var id int64
	id = 1

	product := models.Product{
		ID:    1,
		Name:  "Product A",
		Price: 1000,
	}

	orderProducts := []models.OrderProduct{
		{
			Product:      product,
			ProductPrice: product.Price,
			Amount:       6,
			Total:        6000,
		},
	}

	order := models.Order{
		ID:            1,
		Total:         6000,
		OrderProducts: orderProducts,
	}

	repo.Mock.On("Get", id).Return(order, nil)

	res, err := service.Get(id)
	assert.Equal(t, nil, err)
	assert.Equal(t, id, res.ID)
}

func TestGetFail(t *testing.T) {
	repo, service := Init()

	var id int64
	id = 1

	order := models.Order{}

	errReturn := errors.New("get error")
	repo.Mock.On("Get", id).Return(order, errReturn)

	res, err := service.Get(id)
	assert.Equal(t, errReturn, err)
	assert.NotEqual(t, id, res.ID)
}
