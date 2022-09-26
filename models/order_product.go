package models

type OrderProduct struct {
	ID           int64
	Product      Product
	Order        Order
	ProductPrice int64
	Amount       int64
	Total        int64
}
