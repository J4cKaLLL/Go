package http

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"

	usecase "corp/fif/inte/customers/internal/core/usecase"
)

var (
	uri     = "/falabella"
	version = "/v1"
)

// Server ...
type Server struct {
	CustomerUseCase usecase.CustomerUseCase
}

// NewHandlerServer ...
func NewHandlerServer(e *echo.Echo, customerUseCase usecase.CustomerUseCase) *Server {
	handler := &Server{
		CustomerUseCase: customerUseCase,
	}
	e.POST(uri+version+"/", handler.FindCustomer)
	return handler
}

// FindCustomer ...
func (s *Server) FindCustomer(c echo.Context) error {
	documentNumber := c.FormValue("document_number")
	documentType := c.FormValue("document_type")
	country := c.FormValue("country")

	if len(documentNumber) == 0 || len(documentType) == 0 || len(country) == 0 {
		return c.JSON(http.StatusBadRequest, errors.New("Request Invalida"))
	}

	//Implementa caso de uso (FindCustomer)
	cu, err := s.CustomerUseCase.FindCustomer(documentNumber, documentType, country)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, cu)
}
