package accounts

import "time"

type Account struct {
	Id         string    `json:"id" bson:"_id"`
	Name       string    `json:"name" bson:"name"`
	LastName   string    `json:"last_name" bson:"last_name"`
	MotherName string    `json:"mother_name" bson:"mother_name"`
	Document   string    `json:"document" bson:"document"`
	BirthDate  string    `json:"birth_date" bson:"birth_date"`
	Address    Address   `json:"address" bson:"address"`
	Phone      Phone     `json:"phone" bson:"phone"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
}

type Address struct {
	ZipCode string `json:"zip_code" bson:"zip_code"`
	Number  string `json:"number" bson:"number"`
}

type Phone struct {
	Code   string `json:"code" bson:"code"`
	Number string `json:"number" bson:"number"`
}
