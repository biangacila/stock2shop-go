package domain


type Params struct {
	Key  string
	Val  string
	Type string
}
type Generate struct {
	Id      string
	Org     string
	AppName string
	Ref     float64

	Profile     map[string]interface{}
	OrgDateTime string
	Date        string
	Time        string
	Status      string
}

//let table global
type Util struct {
	AppName     string
	Id          string
	Org         string
	Module      string
	Category    string
	Name        string
	Description string
	Color       string
	Position    float64

	Profile     map[string]interface{}
	OrgDateTime string
	Date        string
	Time        string
	Status      string
}
type Roles struct {
	AppName string
	Id      string
	Org     string
	Module  string

	Name        string
	Description string

	Profile     map[string]interface{}
	OrgDateTime string
	Date        string
	Time        string
	Status      string
}
type Company struct {
	Id             string
	AppName        string
	CustomerNumber string
	Name           string

	ContactEmail  string
	ContactNumber string
	ContactName   string

	//Basic information
	DisplayName             string
	LegalTradingName        string
	Logo                    string
	LineOfBusiness          string
	OrganisationType        string
	BusinessRegistration    string
	OrganisationDescription string

	//Contact Details
	PhysicalSameAsPostal  string
	PhysicalStreetAddress string
	PhysicalCity          string
	PhysicalState         string
	PhysicalPostalCode    string
	PhysicalCountry       string
	PhysicalAttention     string

	PostalStreetAddress string
	PostalCity          string
	PostalState         string
	PostalPostalCode    string
	PostalCountry       string
	PostalAttention     string

	Telephone string
	Email     string
	Website   string

	Profile     map[string]interface{}
	OrgDateTime string
	Date        string
	Time        string
	Status      string
}
type User struct {
	AppName  string
	Id       string
	Org      string
	Username string
	Password string
	Role     string
	Name     string
	Surname  string
	FullName string
	Phone    string
	Email    string
	Account  string

	Profile     map[string]interface{}
	OrgDateTime string
	Date        string
	Time        string
	Status      string
}

