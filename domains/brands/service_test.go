package brands

//
//import (
//	"errors"
//	"testing"
//	"tirenn/test-efishery/fetch-app/models"
//
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//)
//
//func Init() (*RepositoryMock, Service) {
//	var repo = &RepositoryMock{Mock: mock.Mock{}}
//	return repo, Service{repository: repo}
//}
//
//func TestGetSuccess(t *testing.T) {
//	repo, service := Init()
//
//	komoditas := []models.Komoditas{
//		{
//			UUID:         "uuid-1",
//			Komoditas:    "komoditas-1",
//			AreaProvinsi: "area-provinsi-1",
//			AreaKota:     "area-kota-1",
//			Size:         "size-1",
//			Price:        "price-1",
//			TglParsed:    "tgl-parsed-1",
//			Timestamp:    "timestamp-1",
//		},
//		{
//			UUID:         "uuid-2",
//			Komoditas:    "komoditas-2",
//			AreaProvinsi: "area-provinsi-2",
//			AreaKota:     "area-kota-2",
//			Size:         "size-2",
//			Price:        "price-2",
//			TglParsed:    "tgl-parsed-2",
//			Timestamp:    "timestamp-2",
//		},
//	}
//	repo.Mock.On("Get").Return(komoditas, nil)
//
//	res, err := service.Get()
//	assert.Equal(t, nil, err)
//	assert.Equal(t, komoditas[0].UUID, res[0].UUID)
//}
//
//func TestGetFail(t *testing.T) {
//	repo, service := Init()
//
//	komoditas := []models.Komoditas{}
//
//	err := errors.New("Get Error")
//	repo.Mock.On("Get").Return(komoditas, err)
//
//	res, errRes := service.Get()
//	assert.Equal(t, err, errRes)
//	assert.Empty(t, res)
//}
//
//func TestGetAggregateSuccess(t *testing.T) {
//	repo, service := Init()
//
//	komoditas := []models.Komoditas{
//		{
//			UUID:         "uuid-1",
//			Komoditas:    "komoditas-1",
//			AreaProvinsi: "area-provinsi-1",
//			AreaKota:     "area-kota-1",
//			Size:         "size-1",
//			Price:        "price-1",
//			TglParsed:    "tgl-parsed-1",
//			Timestamp:    "timestamp-1",
//		},
//		{
//			UUID:         "uuid-2",
//			Komoditas:    "komoditas-2",
//			AreaProvinsi: "area-provinsi-2",
//			AreaKota:     "area-kota-2",
//			Size:         "size-2",
//			Price:        "price-2",
//			TglParsed:    "tgl-parsed-2",
//			Timestamp:    "timestamp-2",
//		},
//	}
//	repo.Mock.On("Get").Return(komoditas, nil)
//
//	res, err := service.GetAggregate()
//	assert.Equal(t, nil, err)
//	assert.Equal(t, komoditas[0].AreaProvinsi, res[0].AreaProvinsi)
//}
//
//func TestGetAggregateFail(t *testing.T) {
//	repo, service := Init()
//
//	komoditas := []models.Komoditas{}
//
//	err := errors.New("Get Error")
//	repo.Mock.On("Get").Return(komoditas, err)
//
//	res, errRes := service.GetAggregate()
//	assert.Equal(t, err, errRes)
//	assert.Empty(t, res)
//}
