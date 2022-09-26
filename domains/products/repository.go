package products

import (
	"database/sql"
	"tirenn/catalyst/models"
)

type RepositoryContract interface {
	Create(product *models.Product) (err error)
	Get(id int64) (product models.Product, err error)
	GetByBrand(brandID int64) (product []models.Product, err error)
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) RepositoryContract {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(product *models.Product) (err error) {
	stmt, err := r.db.Prepare("INSERT INTO products (id, name, price, brand_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		return
	}

	result, err := stmt.Exec(nil, product.Name, product.Price, product.Brand.ID)
	if err != nil {
		return
	}

	defer stmt.Close()

	lastID, err := result.LastInsertId()
	if err != nil {
		return
	}

	product.ID = lastID
	return
}

func (r *Repository) Get(id int64) (product models.Product, err error) {
	err = r.db.QueryRow("SELECT p.id, p.name, p.price, b.id, b.name FROM products p JOIN brands b ON p.brand_id = b.id WHERE p.id = ?", id).
		Scan(&product.ID, &product.Name, &product.Price, &product.Brand.ID, &product.Brand.Name)
	if err != nil {
		return
	}
	return
}

func (r *Repository) GetByBrand(brandID int64) (products []models.Product, err error) {
	rows, err := r.db.Query("SELECT id, name, price, brand_id FROM products WHERE brand_id = ?", brandID)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		product := models.Product{}
		err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.Brand.ID)
		if err != nil {
			return
		}
		products = append(products, product)
	}

	err = rows.Err()
	if err != nil {
		return
	}

	return
}
