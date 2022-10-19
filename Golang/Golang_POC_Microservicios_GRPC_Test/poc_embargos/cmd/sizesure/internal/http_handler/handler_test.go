package http_handler

import (
	"errors"
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	ent "sizesure/internal/entity"
	mocks "sizesure/internal/usecase/mocks"
)

var (
	// La variable mockAddSeizuresUC es una estructura de arreglos que contiene las posibles combinaciones de request soportados para ser recibidos por el mock
	mockCreateSeizuresUC = []struct {
		inSeize  ent.Seizure
		outSeize  ent.SeizureResp
		outError error
	}{
		{
			ent.Seizure{
				Customer: ent.Client{
					FirstName:   "Eder",
					MiddleName:  "Leandro",
					Surname:     "Rubiano",
					LastSurname: "Camelo",
					Document: ent.DocumentInfo{
						DocumentType:   "Cedula",
						DocumentNumber: "1022331250",
					},
				},
				SeizureInformation: ent.SeizeInfo{
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
					Plaintiff: ent.Applicant{
						FirstName:   "William",
						MiddleName:  "Antonio",
						Surname:     "Buitrago",
						LastSurname: "Herrera",
						Document: ent.DocumentApplicantInfo{
							DocumentType:   "Cedula",
							DocumentNumber: "123456789",
						},
					},
				},
				Product: ent.ProductName{
					ProductID:    "1",
					ProductState: "Activo",
				},
			},
			ent.SeizureResp{
				Codigo: "1",
				Error: "0",
				Mensaje: "Ejecución exitosa",
			},
			nil,
		},
		{
			ent.Seizure{
				Customer: ent.Client{
					FirstName:   "Angela",
					MiddleName:  "Janneth",
					Surname:     "Mora",
					LastSurname: "Gomez",
					Document: ent.DocumentInfo{
						DocumentType:   "Cedula",
						DocumentNumber: "70789456",
					},
				},
				SeizureInformation: ent.SeizeInfo{
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
					Plaintiff: ent.Applicant{
						FirstName:   "Rafael",
						MiddleName:  "Antonio",
						Surname:     "Letrado",
						LastSurname: "Ayala",
						Document: ent.DocumentApplicantInfo{
							DocumentType:   "Cedula",
							DocumentNumber: "987654321",
						},
					},
				},
				Product: ent.ProductName{
					ProductID:    "1",
					ProductState: "Activo",
				},
			},
			ent.SeizureResp{
				Codigo: "",
				Error: "1",
				Mensaje: "Embargo duplicado",
			},			
			errors.New("TestName: Orden de embargo esta duplicada"),
		},
		{
			ent.Seizure{
				Customer: ent.Client{
					FirstName:   "Jonatan",
					MiddleName:  "David",
					Surname:     "Cardenas",
					LastSurname: "Sanchez",
					Document: ent.DocumentInfo{
						DocumentType:   "Cedula",
						DocumentNumber: "70789458",
					},
				},
				SeizureInformation: ent.SeizeInfo{
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
					Plaintiff: ent.Applicant{
						FirstName:   "Juan",
						MiddleName:  "Felipe",
						Surname:     "Castellanos",
						LastSurname: "Ayala",
						Document: ent.DocumentApplicantInfo{
							DocumentType:   "Cedula",
							DocumentNumber: "1234123412",
						},
					},
				},
				Product: ent.ProductName{
					ProductID:    "1",
					ProductState: "Activo",
				},
			},
			ent.SeizureResp{
				Codigo: "",
				Error: "1",
				Mensaje: "Embargo imbalido",
			},
			errors.New("TestName: Orden de Embargo invalida"),
		},
		{
			ent.Seizure{
				Customer: ent.Client{
					FirstName:   "Falabella",
					MiddleName:  "Leandro",
					Surname:     "Rubiano",
					LastSurname: "Camelo",
					Document: ent.DocumentInfo{
						DocumentType:   "Cedula",
						DocumentNumber: "1022331250",
					},
				},
				SeizureInformation: ent.SeizeInfo{
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
					Plaintiff: ent.Applicant{
						FirstName:   "William",
						MiddleName:  "Antonio",
						Surname:     "Buitrago",
						LastSurname: "Herrera",
						Document: ent.DocumentApplicantInfo{
							DocumentType:   "Cedula",
							DocumentNumber: "123456789",
						},
					},
				},
				Product: ent.ProductName{
					ProductID:    "1",
					ProductState: "Activo",
				},
			},
			ent.SeizureResp{
				Codigo: "",
				Error: "1",
				Mensaje: "Error de conexión",
			},
			errors.New("TestName: Prueba de Internal Server Error"),
		},
		{
			ent.Seizure{},
			ent.SeizureResp{
				Codigo: "",
				Error: "1",
				Mensaje: "Error",
			},
			errors.New("TestName: Prueba de request invalida"),
		},
	}

	// Variable expectAddSeizure es una estructura de arreglos que contiene las posibles combinaciones del request que se van a enviar al servicio a modo de peticion
	tests = []struct {
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


func TestHandler_AddSizesures(t *testing.T) {
	e := echo.New()
	e.Validator = NewValidator()

	//crea un nuevo mock con la interfaz definida en el usecase
	mockUseCase := new(mocks.SizesuresUseCase)

	for _, m := range mockCreateSeizuresUC {
		//define el mock que metodo  debe ejecutar y la dependencia
		mockUseCase.On("AddSizesures", m.inSeize).Return(m.outSeize, m.outError)
	}


	for _, tt := range tests {
		t.Run(tt.testcase, func(t *testing.T) {
			//inicializa un handler con el mock del usecase creado
			h := NewServerHandler(e, mockUseCase)
			//incia el request del http (metodo,  URI del metodo, request)
			req := httptest.NewRequest(echo.POST, "/example/embargos/add", strings.NewReader(tt.in))
			//define el header del request http
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			//define el retorno del metodo es decir el response
			rec := httptest.NewRecorder()
			//incializa el context del request y response para invocar el metodo del handler
			c := e.NewContext(req, rec)
			h.AddSizesures(c)
			//valida si la respuesta de la invocación corresponde a la definida en la estructura expectCreateSeizure

			//t.Log("REQUEST:" + tC.in)
			assert.Equal(t, tt.want, rec.Code)

		})
	}
}
