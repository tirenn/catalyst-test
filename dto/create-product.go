package dto

import "tirenn/catalyst/models"

type CreateProduct struct {
	Name    string `json:"name" validate:"required"`
	Price   int64  `json:"price" validate:"numeric"`
	BrandID int64  `json:"brand_id" validate:"numeric"`
}

func ToProduct(createProduct CreateProduct) (product models.Product) {
	product.Price = createProduct.Price
	product.Name = createProduct.Name
	product.BrandID = createProduct.BrandID
	return
}
