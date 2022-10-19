package http

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo"

	beers "example/beers/internal/beers"
)

var (
	errConflict = errors.New("Esta cerveza esta duplicada")
)

// Handler estructura que tiene las dependencias de los Handler
type Handler struct {
	usecase beers.BeerUseCase
}

// NewServerHandler cargando dependencias de caso de uso
func NewServerHandler(e *echo.Echo, usecase beers.BeerUseCase) *Handler {
	h := &Handler{
		usecase: usecase,
	}
	h.RegisterURI(e)
	return h
}

// RegisterURI Creando grupos de URI
func (h *Handler) RegisterURI(e *echo.Echo) {
	o := e.Group("/example/workshop")
	o.GET("/beers", h.SearchBeers)
	o.POST("/beers", h.AddBeers)
}

// SearchBeers handler que expone el metodo GET /beers
func (h *Handler) SearchBeers(c echo.Context) error {
	beers, err := h.usecase.SearchBeers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, beers)
}

// AddBeers ...
func (h *Handler) AddBeers(c echo.Context) error {
	req := new(BeersRequest)
	if err := req.Bind(c); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	beer := beers.Beer{
		ID:       req.ID,
		Name:     req.Name,
		Brewery:  req.Brewery,
		Country:  req.Country,
		Price:    req.Price,
		Currency: req.Currency,
	}
	err := h.usecase.AddBeers(beer)
	log.Print("err")
	log.Print(err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, nil)
}

// BeersRequest ...
type BeersRequest struct {
	ID       int    `json:"ID" validate:"required"`
	Name     string `json:"Name" validate:"required"`
	Brewery  string `json:"Brewery" validate:"required"`
	Country  string `json:"Country" validate:"required"`
	Price    string `json:"Price" validate:"required"`
	Currency string `json:"Currency" validate:"required"`
}

// Bind metodo que valida la request en base a la estructura de BeersRequest
func (r *BeersRequest) Bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}
