package configs

import (
	"database/sql"
	"net/http"
	domain "tirenn/catalyst/domains"
)

func Router(mux *http.ServeMux, db *sql.DB) *http.ServeMux {
	products := domain.InitProductsAPI(db)
	brands := domain.InitBrandsAPI(db)

	mux.HandleFunc("/products", products.ServeHTTP)

	mux.HandleFunc("/brands", brands.ServeHTTP)

	return mux
}
