package model


/*CustomerPhones model response*/
type CustomerPhones struct {
	CustomerName string `json:"customer_name"`
	Country      string `json:"country"`
	CountryCode  string `json:"country_code"`
	PhoneNumber  string `json:"phone_number"`
	State        string `json:"state"`
}

/* Possible state for phone number Validation*/
const (
	NumberOKState    = "ok"
	NumberNotOKState = "Nok"
)
