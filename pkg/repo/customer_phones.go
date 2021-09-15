package repo

import model "jumia-task/pkg/models"

/*CustomerPhonesRepo Holds everything related to the response pattern*/
type CustomerPhonesRepo interface {
	GetCustomerPhoneDetails(info model.CustomerInfo, phone model.Phone, valid bool, number string) (customerPhone model.CustomerPhones)
}

func NewCustomerPhonesRepo() CustomerPhonesRepo {
	return &customerPhonesImp{}
}

type customerPhonesImp struct {
}

/*GetCustomerPhoneDetails for response*/
func (c customerPhonesImp) GetCustomerPhoneDetails(info model.CustomerInfo, phone model.Phone, valid bool, number string) (customerPhone model.CustomerPhones) {
	if valid {
		customerPhone.State = model.NumberOKState
	} else {
		customerPhone.State = model.NumberNotOKState
	}
	customerPhone.CustomerName = info.Name
	customerPhone.CountryCode = phone.Code
	customerPhone.Country = phone.Country
	customerPhone.PhoneNumber = number

	return customerPhone
}
