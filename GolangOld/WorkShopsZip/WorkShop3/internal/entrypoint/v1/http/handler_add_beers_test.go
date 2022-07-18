package http_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	beers "example/beers/internal/beers"
	mocks "example/beers/internal/beers/mocks"
	entrypoint "example/beers/internal/entrypoint/v1/http"
	middleware "example/beers/internal/infrastructure/server/http/middleware"
)

var (
	mockAddBeersUC = []struct {
		inBeers  beers.Beer
		outError error
	}{
		{
			beers.Beer{
				ID:       1,
				Name:     "Corona",
				Brewery:  "Modelo",
				Country:  "Mexico",
				Price:    "5",
				Currency: "USD",
			},
			nil,
		},
		{
			beers.Beer{
				ID:       3,
				Name:     "Torobayo",
				Brewery:  "Kunstman",
				Country:  "Chile",
				Price:    "6",
				Currency: "USD",
			},
			errors.New("Esta cerveza esta duplicada"),
		},
		{
			beers.Beer{
				ID:       100,
				Name:     "Cocacola",
				Brewery:  "Cocacola Company",
				Country:  "USA",
				Price:    "1",
				Currency: "USD",
			},
			errors.New("Esto no es una cerveza"),
		},
	}

	expectAddBear = []struct {
		testcase string
		in       string
		want     int
	}{
		{
			"Prueba de request valida y exitosa",
			`{"ID":1,"Name":"Corona","Brewery":"Modelo","Country":"Mexico","Price":"5","Currency":"USD"}`,
			http.StatusCreated,
		},
		{
			"Prueba de cerveza duplicada en BD",
			`{"ID":3,"Name":"Torobayo","Brewery":"Kunstman","Country":"Chile","Price":"6","Currency":"USD"}`,
			http.StatusConflict,
		},
		{
			"Prueba de request sin información",
			`{}`,
			http.StatusBadRequest,
		},
		{
			"Prueba de request invalida",
			`{"ID":1,"nombre":"Corona","marca":"Modelo","país":"Mexico","Price":"5","Currency":"USD"}`,
			http.StatusBadRequest,
		},
		{
			"Prueba de Internal Server Error",
			`{"ID":100,"Name":"Cocacola","Brewery":"Cocacola Company","Country":"USA","Price":"1","Currency":"USD"}`,
			http.StatusInternalServerError,
		},
	}
)

func TestAddBeers(t *testing.T) {

	e := echo.New()
	e.Validator = middleware.NewValidator()

	mockUseCase := new(mocks.BeerUseCase)

	for _, m := range mockAddBeersUC {
		mockUseCase.On("AddBeers", m.inBeers).Return(m.outError)
	}

	for _, tt := range expectAddBear {

		t.Run(tt.testcase, func(t *testing.T) {
			h := entrypoint.NewServerHandler(e, mockUseCase)
			req := httptest.NewRequest(echo.POST, "/example/workshop/beers", strings.NewReader(tt.in))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			h.AddBeers(c)
			t.Logf("REQUEST:" + tt.in)
			assert.Equal(t, tt.want, rec.Code)

		})
	}

}
