package brands

import (
	"database/sql"
	"log"
	"tirenn/catalyst/models"
)

type RepositoryContract interface {
	Create(brand *models.Brand) (err error)
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) RepositoryContract {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(brand *models.Brand) (err error) {
	log.Print(brand.Name)
	stmt, err := r.db.Prepare("INSERT INTO brands (id, name) VALUES (?, ?)")
	if err != nil {
		return
	}

	result, err := stmt.Exec(nil, brand.Name)
	if err != nil {
		return
	}

	defer stmt.Close()

	lastID, err := result.LastInsertId()
	if err != nil {
		return
	}

	brand.ID = lastID
	return
}
