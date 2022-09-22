package products

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
	get(w http.ResponseWriter, r *http.Request)
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
	case http.MethodGet:
		if r.URL.Query().Get("id") != "" {
			c.get(w, r)
			break
		} else if r.URL.Query().Get("brand_id") != "" {
			c.getByBrand(w, r)
			break
		}
		utils.NotFound(w, errors.New("not found"))
	default:
		utils.NotFound(w, errors.New("method not found"))
	}
	return
}

func (c *Controller) create(w http.ResponseWriter, r *http.Request) {
	var createProduct dto.CreateProduct
	if err := json.NewDecoder(r.Body).Decode(&createProduct); err != nil {
		utils.InternalServerError(w, err)
		return
	}

	if err := utils.ValidateStruct(createProduct); err != nil {
		utils.BadRequest(w, err)
		return
	}

	product, err := c.service.Create(createProduct)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	res, err := json.Marshal(product)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	utils.Created(w, res)
	return
}

func (c *Controller) get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.BadRequest(w, errors.New("id is empty"))
		return
	}

	productID := utils.StringToInt64(id, 0)
	product, err := c.service.Get(productID)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	res, err := json.Marshal(product)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	utils.OK(w, res)
	return
}

func (c *Controller) getByBrand(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("brand_id")
	if id == "" {
		utils.BadRequest(w, errors.New("brand_id is empty"))
		return
	}

	brandID := utils.StringToInt64(id, 0)
	products, err := c.service.GetByBrand(brandID)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	res, err := json.Marshal(products)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	utils.OK(w, res)
	return
}
