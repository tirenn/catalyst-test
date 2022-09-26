package dto

import (
	"tirenn/catalyst/models"
)

type Product struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
	Brand *Brand `json:"brand,omitempty"`
}

type CreateProduct struct {
	Name    string `json:"name" validate:"required"`
	Price   int64  `json:"price" validate:"required"`
	BrandID int64  `json:"brand_id" validate:"required"`
}

func ToProductDTO(product models.Product, brand models.Brand) (productDTO Product) {
	productDTO.ID = product.ID
	productDTO.Name = product.Name
	productDTO.Price = product.Price
	if brand.ID != 0 {
		brandDTO := ToBrandDTO(brand, nil)
		productDTO.Brand = &brandDTO
	}

	return
}

func ToProductsDTO(products []models.Product) (productsDTO []Product) {
	if len(products) == 0 {
		productsDTO = []Product{}
	}

	for _, p := range products {
		productDTO := ToProductDTO(p, models.Brand{})
		productsDTO = append(productsDTO, productDTO)
	}

	return
}

func ToProduct(createProduct CreateProduct) (product models.Product) {
	product.Price = createProduct.Price
	product.Name = createProduct.Name
	product.Brand.ID = createProduct.BrandID
	return
}
