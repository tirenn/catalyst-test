package dto

import (
	"tirenn/catalyst/models"
)

type Brand struct {
	ID       int64      `json:"id"`
	Name     string     `json:"name"`
	Products *[]Product `json:"products,omitempty"`
}

func ToBrandDTO(brand models.Brand, products []Product) (brandDTO Brand) {
	brandDTO.ID = brand.ID
	brandDTO.Name = brand.Name
	brandDTO.Products = &products
	return
}
