package orders

import (
	"context"
	"database/sql"
	"tirenn/catalyst/models"
)

type RepositoryContract interface {
	Create(order *models.Order) (err error)
	GetProduct(id int64) (product models.Product, err error)
	Get(id int64) (order models.Order, err error)
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) RepositoryContract {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(order *models.Order) (err error) {
	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	result, err := tx.Exec(
		"INSERT INTO orders (id, total) VALUES (?, ?)",
		nil, order.Total)
	if err != nil {
		tx.Rollback()
		return err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	order.ID = lastID

	for _, orderProduct := range order.OrderProducts {
		result, err := tx.Exec(
			"INSERT INTO orders_products (id, product_id, order_id, product_price, amount, total) VALUES (?, ?, ?, ?, ?, ?)",
			nil, orderProduct.Product.ID, order.ID, orderProduct.ProductPrice, orderProduct.Amount, orderProduct.Total)
		if err != nil {
			tx.Rollback()
			return err
		}

		lastID, err := result.LastInsertId()
		if err != nil {
			tx.Rollback()
			return err
		}

		orderProduct.ID = lastID
	}

	err = tx.Commit()
	if err != nil {
		return
	}

	return
}

func (r *Repository) GetProduct(id int64) (product models.Product, err error) {
	err = r.db.QueryRow("SELECT id, name, price FROM products WHERE id = ?", id).
		Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return
	}
	return
}

func (r *Repository) Get(id int64) (order models.Order, err error) {
	rows, err := r.db.Query(`
		SELECT op.id, op.product_price, op.amount, op.total,
		       p.id, p.name, p.price,
		       o.id, o.total
		FROM orders_products op 
		    JOIN orders o on o.id = op.order_id
			JOIN products p on p.id = op.product_id
		WHERE op.order_id = ?`, id)
	if err != nil {
		return
	}

	defer rows.Close()

	var orderProducts []models.OrderProduct
	for rows.Next() {
		product := models.Product{}
		orderProduct := models.OrderProduct{}
		err = rows.Scan(
			&orderProduct.ID, &orderProduct.ProductPrice, &orderProduct.Amount, &orderProduct.Total,
			&product.ID, &product.Name, &product.Price,
			&order.ID, &order.Total)
		if err != nil {
			return
		}

		orderProduct.Product = product
		orderProducts = append(orderProducts, orderProduct)
	}

	order.OrderProducts = orderProducts

	err = rows.Err()
	if err != nil {
		return
	}

	return
}
