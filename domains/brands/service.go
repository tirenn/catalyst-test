package brands

import (
	"tirenn/catalyst/dto"
)

type ServiceContract interface {
	Create(createBrand dto.CreateBrand) (res dto.Brand, err error)
}

type Service struct {
	repository RepositoryContract
}

func NewService(repo RepositoryContract) ServiceContract {
	return &Service{
		repository: repo,
	}
}

func (s *Service) Create(createBrand dto.CreateBrand) (res dto.Brand, err error) {
	brand := dto.ToBrand(createBrand)
	err = s.repository.Create(&brand)
	if err != nil {
		return
	}

	res = dto.ToBrandDTO(brand, nil)
	return
}
