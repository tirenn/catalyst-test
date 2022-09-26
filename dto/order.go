package dto

import "tirenn/catalyst/models"

type Order struct {
	ID            int64           `json:"id"`
	OrderProducts []*OrderProduct `json:"products"`
	Total         int64           `json:"total"`
}

type OrderProduct struct {
	Product      *Product `json:"product"`
	ProductPrice int64    `json:"product_price"`
	Amount       int64    `json:"amount"`
	Total        int64    `json:"total"`
}

type CreateOrder struct {
	OrderProducts []CreateOrderProduct `json:"order_products" validate:"min=1,duplicate_product,required,dive"`
}

type CreateOrderProduct struct {
	ProductID int64 `json:"product_id" validate:"required"`
	Amount    int64 `json:"amount" validate:"required"`
}

func ToOrder(total int64) (order models.Order) {
	order.Total = total
	return
}

func ToOrderDTO(order models.Order) (orderDTO Order) {
	orderDTO.ID = order.ID
	orderDTO.Total = order.Total
	orderDTO.OrderProducts = ToOrderProductsDTO(order.OrderProducts)
	return
}

func ToOrderProductDTO(orderProduct models.OrderProduct) (orderProductDTO OrderProduct) {
	orderProductDTO.ProductPrice = orderProduct.ProductPrice
	orderProductDTO.Amount = orderProduct.Amount
	orderProductDTO.Total = orderProduct.Total
	productDTO := ToProductDTO(orderProduct.Product, models.Brand{})
	orderProductDTO.Product = &productDTO
	return
}

func ToOrderProductsDTO(orderProducts []models.OrderProduct) (orderProductsDTO []*OrderProduct) {
	for _, op := range orderProducts {
		orderProductDTO := ToOrderProductDTO(op)
		orderProductsDTO = append(orderProductsDTO, &orderProductDTO)
	}
	return
}
