package http_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	entrypoint "seizures/internal/entrypoint/v1/http"
	middleware "seizures/internal/infrastructure/server/http/middleware"
	seizures "seizures/internal/seizures"
	mocks "seizures/internal/seizures/mocks"
)

var (
	// La variable mockCreateSeizuresUC es una estructura de arreglos que contiene las posibles combinaciones de request soportados para ser recibidos por el mock
	mockCreateSeizuresUC = []struct {
		inSeize  seizures.Seizure
		outError error
	}{
		{
			seizures.Seizure{
				Customer: seizures.Client{
					FirstName:   "Eder",
					MiddleName:  "Leandro",
					Surname:     "Rubiano",
					LastSurname: "Camelo",
					Document: seizures.DocumentInfo{
						DocumentType:   "Cedula",
						DocumentNumber: "1022331250",
					},
				},
				SeizureInformation: seizures.SeizeInfo{
					Concept:         1,
					SeizureType:     2,
					DebitForm:       3,
					ReceptionDate:   "04-12-2019",
					ProcessNumber:   "10",
					Amount:          1000000,
					City:            "Bogota",
					Address:         "Carrera 4 # 4-44",
					JudicialAccount: "1",
					Entity:          "Banco",
					SeizureDate:     "15-12-2019",
					Plaintiff: seizures.Applicant{
						FirstName:   "William",
						MiddleName:  "Antonio",
						Surname:     "Buitrago",
						LastSurname: "Herrera",
						Document: seizures.DocumentApplicantInfo{
							DocumentType:   "Cedula",
							DocumentNumber: "123456789",
						},
					},
				},
				Product: seizures.ProductName{
					ProductID:    "1",
					ProductState: "Activo",
				},
			},
			nil,
		},
		{
			seizures.Seizure{
				Customer: seizures.Client{
					FirstName:   "Angela",
					MiddleName:  "Janneth",
					Surname:     "Mora",
					LastSurname: "Gomez",
					Document: seizures.DocumentInfo{
						DocumentType:   "Cedula",
						DocumentNumber: "70789456",
					},
				},
				SeizureInformation: seizures.SeizeInfo{
					Concept:         1,
					SeizureType:     2,
					DebitForm:       3,
					ReceptionDate:   "04-12-2019",
					ProcessNumber:   "11",
					Amount:          1000000,
					City:            "Bogota",
					Address:         "Carrera 5 # 5-55",
					JudicialAccount: "2",
					Entity:          "Banco",
					SeizureDate:     "16-12-2019",
					Plaintiff: seizures.Applicant{
						FirstName:   "Rafael",
						MiddleName:  "Antonio",
						Surname:     "Letrado",
						LastSurname: "Ayala",
						Document: seizures.DocumentApplicantInfo{
							DocumentType:   "Cedula",
							DocumentNumber: "987654321",
						},
					},
				},
				Product: seizures.ProductName{
					ProductID:    "1",
					ProductState: "Activo",
				},
			},
			errors.New("TestName: Orden de embargo esta duplicada"),
		},
		{
			seizures.Seizure{
				Customer: seizures.Client{
					FirstName:   "Jonatan",
					MiddleName:  "David",
					Surname:     "Cardenas",
					LastSurname: "Sanchez",
					Document: seizures.DocumentInfo{
						DocumentType:   "Cedula",
						DocumentNumber: "70789458",
					},
				},
				SeizureInformation: seizures.SeizeInfo{
					Concept:         1,
					SeizureType:     2,
					DebitForm:       3,
					ReceptionDate:   "04-12-2019",
					ProcessNumber:   "12",
					Amount:          1000000,
					City:            "Bogota",
					Address:         "Carrera 6 # 6-66",
					JudicialAccount: "3",
					Entity:          "Banco",
					SeizureDate:     "17-12-2019",
					Plaintiff: seizures.Applicant{
						FirstName:   "Juan",
						MiddleName:  "Felipe",
						Surname:     "Castellanos",
						LastSurname: "Ayala",
						Document: seizures.DocumentApplicantInfo{
							DocumentType:   "Cedula",
							DocumentNumber: "1234123412",
						},
					},
				},
				Product: seizures.ProductName{
					ProductID:    "1",
					ProductState: "Activo",
				},
			},
			errors.New("TestName: Orden de Embargo invalida"),
		},
		{
			seizures.Seizure{
				Customer: seizures.Client{
					FirstName:   "Falabella",
					MiddleName:  "Leandro",
					Surname:     "Rubiano",
					LastSurname: "Camelo",
					Document: seizures.DocumentInfo{
						DocumentType:   "Cedula",
						DocumentNumber: "1022331250",
					},
				},
				SeizureInformation: seizures.SeizeInfo{
					Concept:         1,
					SeizureType:     2,
					DebitForm:       3,
					ReceptionDate:   "04-12-2019",
					ProcessNumber:   "10",
					Amount:          1000000,
					City:            "Bogota",
					Address:         "Carrera 4 # 4-44",
					JudicialAccount: "2",
					Entity:          "Banco",
					SeizureDate:     "15-12-2019",
					Plaintiff: seizures.Applicant{
						FirstName:   "William",
						MiddleName:  "Antonio",
						Surname:     "Buitrago",
						LastSurname: "Herrera",
						Document: seizures.DocumentApplicantInfo{
							DocumentType:   "Cedula",
							DocumentNumber: "123456789",
						},
					},
				},
				Product: seizures.ProductName{
					ProductID:    "1",
					ProductState: "Activo",
				},
			},
			errors.New("TestName: Prueba de Internal Server Error"),
		},
		{
			seizures.Seizure{},
			errors.New("TestName: Prueba de request invalido"),
		},
	}

	// Variable expectCreateSeizure es una estructura de arreglos que contiene las posibles combinaciones del request que se van a enviar al servicio a modo de peticion
	expectCreateSeizure = []struct {
		testcase string
		in       string
		want     int
	}{
		// Caso exitoso de creacion de peticion de embargo
		{
			"TestName: Prueba de request valida y exitosa",
			`{"Customer":{"FirstName":"Eder","MiddleName":"Leandro","Surname":"Rubiano","LastSurname":"Camelo","Document":{"DocumentType":"Cedula","DocumentNumber":"1022331250"}},"SeizureInformation":{"Concept":1,"SeizureType":2,"DebitForm":3,"ReceptionDate":"04-12-2019","ProcessNumber":"10","Amount":1000000,"City":"Bogota","Address":"Carrera 4 # 4-44","JudicialAccount":"1","Entity":"Banco","SeizureDate":"15-12-2019","Plaintiff":{"FirstName":"William","MiddleName":"Antonio","Surname":"Buitrago","LastSurname":"Herrera","Document":{"DocumentType":"Cedula","DocumentNumber":"123456789"}}},"Product":{"ProductID":"1","ProductState":"Activo"}}`,
			http.StatusCreated,
		},
		// Caso de orden de embargo duplicada
		{
			"TestName: Orden de embargo esta duplicada",
			`{"Customer":{"FirstName":"Angela","MiddleName":"Janneth","Surname":"Mora","LastSurname":"Gomez","Document":{"DocumentType":"Cedula","DocumentNumber":"70789456"}},"SeizureInformation":{"Concept":1,"SeizureType":2,"DebitForm":3,"ReceptionDate":"04-12-2019","ProcessNumber":"11","Amount":1000000,"City":"Bogota","Address":"Carrera 5 # 5-55","JudicialAccount":"2","Entity":"Banco","SeizureDate":"16-12-2019","Plaintiff":{"FirstName":"Rafael","MiddleName":"Antonio","Surname":"Letrado","LastSurname":"Ayala","Document":{"DocumentType":"Cedula","DocumentNumber":"987654321"}}},"Product":{"ProductID":"1","ProductState":"Activo"}}`,
			http.StatusConflict,
		},
		// Caso de request vacio
		{
			"TestName: Prueba de request invalido",
			`{}`,
			http.StatusBadRequest,
		},
		// Caso de Request Invalido
		{
			"TestName: Prueba de request invalido",
			`{"Customer":{"FirstName":"Jonatan","MiddleName":"David","Surname":"Cardenas","LastSurname":"Sanchez","Document":{"DocumentType":"Cedula","DocumentNumber":"70789458"}},"SeizureInformation":{"Concept":1,"SeizureType":2,"DebitForm":3,"ReceptionDate":"04-12-2019","ProcessNumber":"12","Amount":1000000,"City":"Bogota","Address":"Carrera 6 # 6-66","JudicialAccount":"3","Entity":"Banco","SeizureDate":"17-12-2019","Plaintiff":{"FirstName":"Juan","MiddleName":"Felipe","Surname":"Castellanos","LastSurname":"Ayala","Document":{"DocumentType":"Cedula","DocumentNumber":"1234123412"}}},"Product":{"ProductID":"1","ProductState":"Activo"}}`,
			http.StatusBadRequest,
		},
		// Caso de error interno : "FirstName:"Falabella"
		{
			"TestName: Prueba de Internal Server Error",
			`{"Customer":{"FirstName":"Falabella","MiddleName":"Leandro","Surname":"Rubiano","LastSurname":"Camelo","Document":{"DocumentType":"Cedula","DocumentNumber":"1022331250"}},"SeizureInformation":{"Concept":1,"SeizureType":2,"DebitForm":3,"ReceptionDate":"04-12-2019","ProcessNumber":"10","Amount":1000000,"City":"Bogota","Address":"Carrera 4 # 4-44","JudicialAccount":"2","Entity":"Banco","SeizureDate":"15-12-2019","Plaintiff":{"FirstName":"William","MiddleName":"Antonio","Surname":"Buitrago","LastSurname":"Herrera","Document":{"DocumentType":"Cedula","DocumentNumber":"123456789"}}},"Product":{"ProductID":"1","ProductState":"Activo"}}`,
			http.StatusInternalServerError,
		},
	}
)

func TestCreateSeizure(t *testing.T) {

	//inicializa libreria echo para poder utilizar servicio http
	e := echo.New()

	e.Validator = middleware.NewValidator()

	//crea un nuevo mock con la interfaz definida en el usecase
	mockUseCase := new(mocks.SeizureUseCase)

	for _, m := range mockCreateSeizuresUC {
		//define el mock que metodo  debe ejecutar y la dependencia
		mockUseCase.On("CreateSeizure", m.inSeize).Return(m.outError)
	}

	for _, tC := range expectCreateSeizure {
		t.Run(tC.testcase, func(t *testing.T) {
			//inicializa un handler con el mock del usecase creado
			h := entrypoint.NewServerHandler(e, mockUseCase)
			//incia el request del http (metodo,  URI del metodo, request)
			req := httptest.NewRequest(echo.POST, "/embargos/prueba/addEmbargos", strings.NewReader(tC.in))
			//define el header del request http
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			//define el retorno del metodo es decir el response
			rec := httptest.NewRecorder()
			//incializa el context del request y response para invocar el metodo del handler
			c := e.NewContext(req, rec)
			h.CreateSeizure(c)
			//valida si la respuesta de la invocación corresponde a la definida en la estructura expectCreateSeizure

			//t.Log("REQUEST:" + tC.in)
			assert.Equal(t, tC.want, rec.Code)
		})
	}
}
