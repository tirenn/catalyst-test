package dto

import "tirenn/catalyst/models"

type CreateBrand struct {
	Name string `json:"name" validate:"required"`
}

func ToBrand(createBrand CreateBrand) (brand models.Brand) {
	brand.Name = createBrand.Name
	return
}
