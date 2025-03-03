package http

//importamos las depencias para los handler
import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo"

	//se crea vinculo a la carpeta del core
	seizures "seizures/internal/seizures"
)

var (
	errDuplicada           = errors.New("TestName: Orden de embargo esta duplicada")
	errInvalida            = errors.New("TestName: Orden de Embargo invalida")
	errBadRequest          = errors.New("TestName: Prueba de request invalido")
	errInternalServerError = errors.New("TestName: Prueba de Internal Server Error")
)

// Handler Estrutura de los handlers heredada del usecase interfaz
type Handler struct {
	usecase seizures.SeizureUseCase
}

// CreateSeizure funcion que sera invocada por cada metodo
func (h *Handler) CreateSeizure(c echo.Context) error {

	req := new(SeizureRequest)
	// Valida si el request SeizureRequest instanciado en la variable 'req' cumple tanto con la estructura definida como con las propiedades de cada elemento
	if err := req.Bind(c); err != nil {
		//log.Print(c)
		return c.JSON(http.StatusBadRequest, err)
	}
	SeizureRequest := seizures.Seizure{
		Customer: seizures.Client{
			FirstName:   req.Customer.FirstName,
			MiddleName:  req.Customer.MiddleName,
			Surname:     req.Customer.Surname,
			LastSurname: req.Customer.LastSurname,
			Document: seizures.DocumentInfo{
				DocumentType:   req.Customer.Document.DocumentType,
				DocumentNumber: req.Customer.Document.DocumentNumber,
			},
		},
		SeizureInformation: seizures.SeizeInfo{
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
			Plaintiff: seizures.Applicant{
				FirstName:   req.SeizureInformation.Plaintiff.FirstName,
				MiddleName:  req.SeizureInformation.Plaintiff.MiddleName,
				Surname:     req.SeizureInformation.Plaintiff.Surname,
				LastSurname: req.SeizureInformation.Plaintiff.LastSurname,
				Document: seizures.DocumentApplicantInfo{
					DocumentType:   req.SeizureInformation.Plaintiff.Document.DocumentType,
					DocumentNumber: req.SeizureInformation.Plaintiff.Document.DocumentNumber,
				},
			},
		},
		Product: seizures.ProductName{
			ProductID:    req.Product.ProductID,
			ProductState: req.Product.ProductState,
		},
	}
	// valida los demas casos de error posibles
	err := h.usecase.CreateSeizure(SeizureRequest)
	if err != nil {

		if err.Error() == errDuplicada.Error() {
			//log.Print("errDuplicada" + err.Error())
			return c.JSON(http.StatusConflict, err)
		}
		if err.Error() == errInternalServerError.Error() {
			log.Print("errInternalServerError" + err.Error())
			return c.JSON(http.StatusInternalServerError, err)
		}
	}
	return c.JSON(http.StatusCreated, nil)
}

// RegisterURI 3. funcion para registrar path y metodos del servicio
func (h *Handler) RegisterURI(e *echo.Echo) {
	//agrupación de los puntos finales PATH
	o := e.Group("/embargos/prueba")
	//registrar los metodos al group creado y asocia funcion a ejecutar
	o.POST("/addEmbargos", h.CreateSeizure)
}

// NewServerHandler que permite cargar dependencias
func NewServerHandler(e *echo.Echo, usecase seizures.SeizureUseCase) *Handler {
	//agrega depencia del usecase al handler
	h := &Handler{
		usecase: usecase,
	}
	//registra los metodos y el path del servicio
	h.RegisterURI(e)
	return h
}

// SeizureRequest estructura que muestra las variables de Embargos
type SeizureRequest struct {
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

// Bind metodo que valida la request en base a la estructura de SeizureRequest
func (r *SeizureRequest) Bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}
