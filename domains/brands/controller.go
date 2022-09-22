package brands

import (
	"encoding/json"
	"errors"
	"net/http"
	"tirenn/catalyst/dto"
	"tirenn/catalyst/utils"
)

type ControllerContract interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	create(w http.ResponseWriter, r *http.Request)
}

type Controller struct {
	service ServiceContract
}

func NewController(service ServiceContract) ControllerContract {
	return &Controller{
		service: service,
	}
}

func (c *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.create(w, r)
	default:
		utils.NotFound(w, errors.New("method not found"))
	}
	return
}

func (c *Controller) create(w http.ResponseWriter, r *http.Request) {
	var createBrand dto.CreateBrand
	if err := json.NewDecoder(r.Body).Decode(&createBrand); err != nil {
		utils.InternalServerError(w, err)
		return
	}

	if err := utils.ValidateStruct(createBrand); err != nil {
		utils.BadRequest(w, err)
		return
	}

	brand, err := c.service.Create(createBrand)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	res, err := json.Marshal(brand)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	utils.Created(w, res)
	return
}
