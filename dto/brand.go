package dto

import (
	"tirenn/catalyst/models"
)

type Brand struct {
	ID       int64      `json:"id"`
	Name     string     `json:"name"`
	Products []*Product `json:"products,omitempty"`
}

type CreateBrand struct {
	Name string `json:"name" validate:"required"`
}

func ToBrand(createBrand CreateBrand) (brand models.Brand) {
	brand.Name = createBrand.Name
	return
}

func ToBrandDTO(brand models.Brand, products []*Product) (brandDTO Brand) {
	brandDTO.ID = brand.ID
	brandDTO.Name = brand.Name
	brandDTO.Products = products
	if products == nil {
		brandDTO.Products = nil
	}
	return
}
