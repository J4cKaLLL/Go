package seizures

// Seizure estructura que muestra las variables de Embargos
type Seizure struct {
	Customer           Client      `json:"customer"`
	SeizureInformation SeizeInfo   `json:"seizureInformation"`
	Product            ProductName `json:"product"`
}

// DocumentInfo tipo y numero de documento
type DocumentInfo struct {
	DocumentType   string `json:"documentType"`
	DocumentNumber string `json:"documentNumber"`
}

// Client Datos de cliente
type Client struct {
	FirstName   string       `json:"firstName"`
	MiddleName  string       `json:"middleName"`
	Surname     string       `json:"surname"`
	LastSurname string       `json:"lastSurname"`
	Document    DocumentInfo `json:"document"`
}

// SeizeInfo Informacion de embargo
type SeizeInfo struct {
	Concept         int    `json:"concept"`
	SeizureType     int    `json:"seizureType"`
	DebitForm       int    `json:"debitForm"`
	ReceptionDate   string `json:"receptionDate"`
	ProcessNumber   string `json:"processNumber"`
	Amount          int    `json:"amount"`
	City            string `json:"city"`
	Address         string `json:"address"`
	JudicialAccount string `json:"judicialAccount"`
	Entity          string `json:"entity"`
	SeizureDate     string `json:"seizureDate"`
	Plaintiff       Client `json:"plaintiff"`
}

// ProductName informacion de producto
type ProductName struct {
	ProductID    string `json:"productID"`
	ProductState string `json:"prodictState"`
}
