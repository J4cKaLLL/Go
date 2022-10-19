package http_handler

import (
    "net/http"

    "github.com/labstack/echo"

	sizesure "sizesure/internal/usecase"
	ent "sizesure/internal/entity"
)

// Handler estructura que tiene las dependencias de los Handler
type Handler struct {
    usecase sizesure.SizesuresUseCase
}

// NewServerHandler creando el nuevo server y cargando dependencias caso de uso
func NewServerHandler(e *echo.Echo, usecase sizesure.SizesuresUseCase) *Handler {
    //agrega al handler la dependenci de usecase
    h := &Handler{
        usecase: usecase,    
    }
    //registrar los metodos del servicio
    h.RegisterURI(e)
    return h
}

// RegisterURI creando los grupos de URI
func (h *Handler) RegisterURI(e *echo.Echo) {
    //agrupacion de puntos finales
    o := e.Group("/example/embargos")
    //agrega el methodo al handler
    o.POST("/add",h.AddSizesures)
}

// AddSizesures operacion a realizar
func (h *Handler) AddSizesures(c echo.Context) error {
    req := new(SizesuresRequest)
	// Valida si el request SeizureRequest instanciado en la variable 'req' cumple tanto con la estructura definida como con las propiedades de cada elemento
	if err := req.Bind(c); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	SeizureRequest := ent.Seizure{
		Customer: ent.Client{
			FirstName:   req.Customer.FirstName,
			MiddleName:  req.Customer.MiddleName,
			Surname:     req.Customer.Surname,
			LastSurname: req.Customer.LastSurname,
			Document: ent.DocumentInfo{
				DocumentType:   req.Customer.Document.DocumentType,
				DocumentNumber: req.Customer.Document.DocumentNumber,
			},
		},
		SeizureInformation: ent.SeizeInfo{
			Concept:         req.SeizureInformation.Concept,
			SeizureType:     req.SeizureInformation.SeizureType,
			DebitForm:       req.SeizureInformation.DebitForm,
			ReceptionDate:   req.SeizureInformation.ReceptionDate,
			ProcessNumber:   req.SeizureInformation.ProcessNumber,
			Amount:          req.SeizureInformation.Amount,
			City:            req.SeizureInformation.City,
			Address:         req.SeizureInformation.Address,
			JudicialAccount: req.SeizureInformation.JudicialAccount,
			Entity:          req.SeizureInformation.Entity,
			SeizureDate:     req.SeizureInformation.SeizureDate,
			Plaintiff: ent.Applicant{
				FirstName:   req.SeizureInformation.Plaintiff.FirstName,
				MiddleName:  req.SeizureInformation.Plaintiff.MiddleName,
				Surname:     req.SeizureInformation.Plaintiff.Surname,
				LastSurname: req.SeizureInformation.Plaintiff.LastSurname,
				Document: ent.DocumentApplicantInfo{
					DocumentType:   req.SeizureInformation.Plaintiff.Document.DocumentType,
					DocumentNumber: req.SeizureInformation.Plaintiff.Document.DocumentNumber,
				},
			},
		},
		Product: ent.ProductName{
			ProductID:    req.Product.ProductID,
			ProductState: req.Product.ProductState,
		},
	}
	// valida los demas casos de error posibles
	resp, err := h.usecase.AddSizesures(SeizureRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, resp)
}

// SizesuresRequest estructura del request
type SizesuresRequest struct {
    Customer           Client      `json:"customer" validate:"required"`
	SeizureInformation SeizeInfo   `json:"seizureInformation" validate:"required"`
	Product            ProductName `json:"product" validate:"required"`
}

// DocumentInfo tipo y numero de documento
type DocumentInfo struct {
	DocumentType   string `json:"documentType" validate:"required"`
	DocumentNumber string `json:"documentNumber" validate:"required"`
}

// DocumentApplicantInfo tipo y numero de documento
type DocumentApplicantInfo struct {
	DocumentType   string `json:"documentType" validate:"required"`
	DocumentNumber string `json:"documentNumber" validate:"required"`
}

// Client Datos de cliente
type Client struct {
	FirstName   string       `json:"firstName" validate:"required"`
	MiddleName  string       `json:"middleName" validate:"required"`
	Surname     string       `json:"surname" validate:"required"`
	LastSurname string       `json:"lastSurname" validate:"required"`
	Document    DocumentInfo `json:"document" validate:"required"`
}

// Applicant Datos de cliente
type Applicant struct {
	FirstName   string                `json:"firstName" validate:"required"`
	MiddleName  string                `json:"middleName" validate:"required"`
	Surname     string                `json:"surname" validate:"required"`
	LastSurname string                `json:"lastSurname" validate:"required"`
	Document    DocumentApplicantInfo `json:"document" validate:"required"`
}

// SeizeInfo Informacion de embargo
type SeizeInfo struct {
	Concept         int       `json:"concept" validate:"required"`
	SeizureType     int       `json:"seizureType" validate:"required"`
	DebitForm       int       `json:"debitForm" validate:"required"`
	ReceptionDate   string    `json:"receptionDate" validate:"required"`
	ProcessNumber   string    `json:"processNumber" validate:"required"`
	Amount          int       `json:"amount" validate:"required"`
	City            string    `json:"city" validate:"required"`
	Address         string    `json:"address" validate:"required"`
	JudicialAccount string    `json:"judicialAccount" validate:"required"`
	Entity          string    `json:"entity" validate:"required"`
	SeizureDate     string    `json:"seizureDate" validate:"required"`
	Plaintiff       Applicant `json:"plaintiff" validate:"required"`
}

// ProductName informacion de producto
type ProductName struct {
	ProductID    string `json:"productID" validate:"required"`
	ProductState string `json:"productState" validate:"required"`
}

// Bind metodo para que valide la estructura y campos del request
func (r *SizesuresRequest) Bind(c echo.Context) error {
    //valida estructura
    if err := c.Bind(r); err != nil {
        return err
    }
    //valida campos 
    if err := c.Validate(r); err != nil {
        return err
    }
    return nil
}
