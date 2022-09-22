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

func ToProductDTO(product models.Product, brand models.Brand) (productDTO Product) {
	productDTO.ID = product.ID
	productDTO.Name = product.Name
	productDTO.Price = product.Price
	if brand != (models.Brand{}) {
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
		productDTO := Product{
			ID:    p.ID,
			Name:  p.Name,
			Price: p.Price,
		}
		productsDTO = append(productsDTO, productDTO)
	}

	return
}
