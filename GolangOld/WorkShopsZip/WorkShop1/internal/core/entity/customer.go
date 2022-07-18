package entity

// Customer ...
type Customer struct {
	Person     Person
	IsEmployee bool
	Address    []*Address
	Telephone  []*Telephone
	Email      []*Email
}

// Person ...
type Person struct {
	Document         Document
	FirstName        string
	MiddleName       string
	AditionalName    string
	LastSurname      string
	AditionalSurname string
	FullName         string
	Names            string
	Gender           string
	BirthDate        string
	Nationality      string
	MaritalStatus    string
}

// Address ...
type Address struct {
	TypeAddress          string
	FullAddress          string
	ZipCode              string
	StreetName           string
	StreetNumber         string
	Floor                string
	Apartment            string
	AditionalInformation string
	Country              Country
	State                State
	City                 City
	Township             Township
	Town                 Town
}

// Country ...
type Country struct {
	ID   string
	Id   string
	Name string
}

// State ...
type State struct {
	ID   string
	Id   string
	Name string
}

// City ...
type City struct {
	ID   string
	Id   string
	Name string
}

// Township ...
type Township struct {
	ID   string
	Id   string
	Name string
}

// Town ...
type Town struct {
	ID   string
	Id   string
	Name string
}

// Telephone ...
type Telephone struct {
	TypeTelephone string
	Number        string
	AreaCode      string
	CountryCode   string
}

// Email ...
type Email struct {
	TypeEmail          string
	PreferredEmailFlag string
	EmailAddress       string
}

// Document ...
type Document struct {
	TypeDocument string
	Number       string
}
