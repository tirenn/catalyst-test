package domain

import (
	"database/sql"
	"tirenn/catalyst/domains/brands"
	"tirenn/catalyst/domains/products"
)

func InitProductsAPI(db *sql.DB) products.ControllerContract {
	repositoryContract := products.NewRepository(db)
	serviceContract := products.NewService(repositoryContract)
	controllerContract := products.NewController(serviceContract)
	return controllerContract
}

func InitBrandsAPI(db *sql.DB) brands.ControllerContract {
	repositoryContract := brands.NewRepository(db)
	serviceContract := brands.NewService(repositoryContract)
	controllerContract := brands.NewController(serviceContract)
	return controllerContract
}
