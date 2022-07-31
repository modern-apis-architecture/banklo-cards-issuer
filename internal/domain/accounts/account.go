package accounts

type Account struct {
	Name       string  `json:"name"`
	LastName   string  `json:"last_name"`
	MotherName string  `json:"mother_name"`
	Document   string  `json:"document"`
	BirthDate  string  `json:"birth_date"`
	Address    Address `json:"address"`
	Phone      Phone   `json:"phone"`
}

type Address struct {
	ZipCode string `json:"zip_code"`
	Number  string `json:"number"`
}

type Phone struct {
	Code   string `json:"code"`
	Number string `json:"number"`
}
