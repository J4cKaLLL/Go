package http

import (
	"errors"
	"net/http"
	seizures "seizure/internal/seizures"
)

var (
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
					Plaintiff: seizures.Client{
						FirstName:   "William",
						MiddleName:  "Antonio",
						Surname:     "Buitrago",
						LastSurname: "Herrera",
						Document: seizures.DocumentInfo{
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
					Plaintiff: seizures.Client{
						FirstName:   "Rafael",
						MiddleName:  "Antonio",
						Surname:     "Letrado",
						LastSurname: "Ayala",
						Document: seizures.DocumentInfo{
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
			errors.New("Esta Orden de embargo esta duplicada"),
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
					Plaintiff: seizures.Client{
						FirstName:   "Juan",
						MiddleName:  "Felipe",
						Surname:     "Castellanos",
						LastSurname: "Ayala",
						Document: seizures.DocumentInfo{
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
			errors.New("Orden de Embargo invalida"),
		},
	}

	expectCreateSeizure = []struct {
		testcase string
		in       string
		want     int
	}{
		{
			"Prueba de request valida y exitosa",
			`{"Seizure":[{"Customer":[{"FirstName":"Eder","MiddleName":"Leandro","Surname":"Rubiano","LastSurname":"Camelo","Document":[{"DocumentType":"Cedula","DocumentNumber":"1022331250"}]}]}],"SeizureInformation":[{"Concept":1,"SeizureType":2,"DebitForm":3,"ReceptionDate":"04-12-2019","ProcessNumber":"10","Amount":1000000,"City":"Bogota","Address":"Carrera 4 # 4-44","JudicialAccount":"1","Entity":"Banco","SeizureDate":"15-12-2019","Plaintiff":[{"FirstName":"William","MiddleName":"Antonio","Surname":"Buitrago","LastSurname":"Herrera","Document":[{"DocumentType":"Cedula","DocumentNumber":"123456789"}]}]}],"ProductName":[{"ProductID":"1","ProductState":"Activo"}]}`,
			http.StatusCreated,
		},
		{
			"Prueba de Embargo duplicado",
			`{"Seizure":[{"Customer":[{"FirstName":"Angela","MiddleName":"Janneth","Surname":"Mora","LastSurname":"Gomez","Document":[{"DocumentType":"Cedula","DocumentNumber":"70789456"}]}]}],"SeizureInformation":[{"Concept":1,"SeizureType":2,"DebitForm":3,"ReceptionDate":"04-12-2019","ProcessNumber":"10","Amount":1000000,"City":"Bogota","Address":"Carrera 5 # 5-55","JudicialAccount":"2","Entity":"Banco","SeizureDate":"16-12-2019","Plaintiff":[{"FirstName":"Rafael","MiddleName":"Antonio","Surname":"Letrado","LastSurname":"Ayala","Document":[{"DocumentType":"Cedula","DocumentNumber":"987654321"}]}]}],"ProductName":[{"ProductID":"1","ProductState":"Activo"}]}`,
			http.StatusConflict,
		},
		{
			"Prueba de request sin información",
			`{}`,
			http.StatusBadRequest,
		},
		{
			"Prueba de request invalida",
			`{"Seizure":[{"Customer":[{"FirstName":"Jonatan","MiddleName":"David","Surname":"Cardenas","LastSurname":"Sanchez","Document":[{"DocumentType":"Cedula","DocumentNumber":"70789458"}]}]}],"SeizureInformation":[{"Concept":1,"SeizureType":2,"DebitForm":3,"ReceptionDate":"04-12-2019","ProcessNumber":"12","Amount":1000000,"City":"Bogota","EmailAddress":"Carrera 6 # 6-66","JudicialAccount":"3","Entity":"Banco","SeizureDate":"17-12-2019","Plaintiff":[{"FirstName":"Juan","MiddleName":"Felipe","Surname":"Castellanos","LastSurname":"Ayala","Document":[{"DocumentType":"Cedula","DocumentNumber":"1234123412"}]}]}],"ProductName":[{"ProductID":"1","ProductState":"Activo"}]}`,
			http.StatusBadRequest,
		},
		{
			"Prueba de Internal Server Error",
			`{"Seizure":[{"Customer":[{"FirstName":"Angela","MiddleName":"Janneth","Surname":"Mora","LastSurname":"Gomez","Document":[{"DocumentType":"Cedula","DocumentNumber":"70789456"}]}]}],"SeizureInformation":[{"Concept":1,"SeizureType":2,"DebitForm":3,"ReceptionDate":"04-12-2019","ProcessNumber":"10","Amount":1000000,"City":"Bogota","Address":"Carrera 5 # 5-55","JudicialAccount":"2","Entity":"Banco","SeizureDate":"16-12-2019","Plaintiff":[{"FirstName":"Rafael","MiddleName":"Antonio","Surname":"Letrado","LastSurname":"Ayala","Document":[{"DocumentType":"Cedula","DocumentNumber":"987654321"}]}]}],"ProductName":[{"ProductID":"Indefinido","ProductState":"Activo"}]}`,
			http.StatusInternalServerError,
		},
	}
)

/*func TestCreateSeizure(t *testing.T) {
	//Ayuda WILLIAM !!!
}*/
