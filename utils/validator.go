package utils

import (
	"github.com/go-playground/validator/v10"
	"tirenn/catalyst/dto"
)

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
	validate.RegisterValidation("duplicate_product", validateDuplicateProducts)
}

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

func validateDuplicateProducts(fl validator.FieldLevel) bool {
	v := fl.Field().Interface().([]dto.CreateOrderProduct)

	visited := make(map[int64]bool, 0)
	for _, co := range v {
		if visited[co.ProductID] == true {
			return false
		} else {
			visited[co.ProductID] = true
		}
	}

	return true
}
