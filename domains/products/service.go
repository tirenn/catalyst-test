package products

import (
	"tirenn/catalyst/dto"
	"tirenn/catalyst/models"
)

type ServiceContract interface {
	Create(createProduct dto.CreateProduct) (res dto.Product, err error)
	Get(id int64) (res dto.Product, err error)
	GetByBrand(brandID int64) (res []dto.Product, err error)
}

type Service struct {
	repository RepositoryContract
}

func NewService(repo RepositoryContract) ServiceContract {
	return &Service{
		repository: repo,
	}
}

func (s *Service) Create(createProduct dto.CreateProduct) (res dto.Product, err error) {
	product := dto.ToProduct(createProduct)
	err = s.repository.Create(&product)
	if err != nil {
		return
	}

	res = dto.ToProductDTO(product, models.Brand{})
	return
}

func (s *Service) Get(id int64) (res dto.Product, err error) {
	product, err := s.repository.Get(id)
	if err != nil {
		return
	}

	brand, err := s.repository.GetBrand(product.BrandID)
	if err != nil {
		return
	}

	res = dto.ToProductDTO(product, brand)
	return
}

func (s *Service) GetByBrand(brandID int64) (res []dto.Product, err error) {
	products, err := s.repository.GetByBrand(brandID)
	if err != nil {
		return
	}

	res = dto.ToProductsDTO(products)
	return
}
