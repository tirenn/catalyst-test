package orders

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
	case http.MethodGet:
		if r.URL.Query().Get("id") != "" {
			c.get(w, r)
			break
		}
		utils.NotFound(w, errors.New("not found"))
	default:
		utils.NotFound(w, errors.New("method not found"))
	}
	return
}

func (c *Controller) create(w http.ResponseWriter, r *http.Request) {
	var createOrder dto.CreateOrder
	if err := json.NewDecoder(r.Body).Decode(&createOrder); err != nil {
		utils.InternalServerError(w, err)
		return
	}

	if err := utils.ValidateStruct(createOrder); err != nil {
		utils.BadRequest(w, err)
		return
	}

	order, err := c.service.Create(createOrder.OrderProducts)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	res, err := json.Marshal(order)
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

	orderID := utils.StringToInt64(id, 0)
	order, err := c.service.Get(orderID)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	res, err := json.Marshal(order)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	utils.OK(w, res)
	return
}
