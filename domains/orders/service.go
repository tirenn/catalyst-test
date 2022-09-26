package orders

import (
	"tirenn/catalyst/dto"
	"tirenn/catalyst/models"
)

type ServiceContract interface {
	Create(createOrder []dto.CreateOrderProduct) (res dto.Order, err error)
	Get(id int64) (res dto.Order, err error)
}

type Service struct {
	repository RepositoryContract
}

func NewService(repo RepositoryContract) ServiceContract {
	return &Service{
		repository: repo,
	}
}

func (s *Service) Create(createOrderProducts []dto.CreateOrderProduct) (res dto.Order, err error) {
	total := int64(0)
	var orderProducts []models.OrderProduct
	for _, c := range createOrderProducts {
		orderProduct := models.OrderProduct{}
		product, err := s.repository.GetProduct(c.ProductID)
		if err != nil {
			return dto.Order{}, err
		}
		orderProduct.Product = product
		orderProduct.ProductPrice = product.Price
		orderProduct.Amount = c.Amount
		orderProduct.Total = product.Price * c.Amount

		orderProducts = append(orderProducts, orderProduct)
		total = total + orderProduct.Total
	}

	order := dto.ToOrder(total)
	order.OrderProducts = orderProducts

	err = s.repository.Create(&order)
	if err != nil {
		return
	}

	res = dto.ToOrderDTO(order)
	return
}

func (s *Service) Get(id int64) (res dto.Order, err error) {
	order, err := s.repository.Get(id)
	if err != nil {
		return
	}

	res = dto.ToOrderDTO(order)
	return
}
