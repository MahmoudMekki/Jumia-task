package model


type CustomerPhones struct{
	CustomerName		string  `json:"customer_name"`
	Country				string	`json:"country"`
	CountryCode			string	`json:"country_code"`
	PhoneNumber			string	`json:"phone_number"`
	State				string	`json:"state"`
}

const(
	NumberOKState = "ok"
	NumberNotOKState = "Nok"
)