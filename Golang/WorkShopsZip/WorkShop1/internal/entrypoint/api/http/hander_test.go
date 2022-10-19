package http_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	usecase "corp/fif/inte/customers/internal/core/usecase/mocks"
	entrypoint "corp/fif/inte/customers/internal/entrypoint/api/http"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestFindCustomer(t *testing.T) {
	// New server echo
	e := echo.New()

	// Crear Mocks UseCase
	mockUseCase := new(usecase.CustomerUseCase)
	mockUseCase.On("FindCustomer", "1", "DNI", "Argentina").Return(nil, nil)
	mockUseCase.On("FindCustomer", "2", "DNI", "Argentina").Return(nil, errors.New("Internal Server Error"))
	mockUseCase.On("FindCustomer", "", "", "").Return(nil, errors.New("Internal Server Error"))

	// Prueba de request valida y con respuesta OK
	f := make(url.Values)
	f.Set("document_number", "1")
	f.Set("document_type", "DNI")
	f.Set("country", "Argentina")
	req := httptest.NewRequest(echo.POST, "/falabella/v1/", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := entrypoint.NewHandlerServer(e, mockUseCase)
	err := h.FindCustomer(c)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Nil(t, err)

	// Prueba de request valida y con respuesta con Errores
	f = make(url.Values)
	f.Set("document_number", "2")
	f.Set("document_type", "DNI")
	f.Set("country", "Argentina")
	req = httptest.NewRequest(echo.POST, "/falabella/v1/", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	h = entrypoint.NewHandlerServer(e, mockUseCase)
	err = h.FindCustomer(c)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)

	// Prueba de request Invalida
	f = make(url.Values)
	f.Set("document_number", "")
	f.Set("document_type", "")
	f.Set("country", "")
	req = httptest.NewRequest(echo.POST, "/falabella/v1/", strings.NewReader(f.Encode()))
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	h = entrypoint.NewHandlerServer(e, mockUseCase)
	h.FindCustomer(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

}
