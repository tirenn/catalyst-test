package models

type OrderProduct struct {
	ID           int64
	ProductID    int64
	OrderID      int64
	ProductPrice int64
	Amount       int64
	Total        int64
}
