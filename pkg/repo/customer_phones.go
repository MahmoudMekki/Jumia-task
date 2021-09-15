package repo

import model "jumia-task/pkg/models"

type CustomerPhonesRepo interface {
	GetCustomerPhoneDetails(info model.CutomerInfo,phone model.Phone,valid bool,number string)(customerPhone model.CustomerPhones)
}

func NewCustomerPhonesRepo() CustomerPhonesRepo {
	return &customerPonesImp{}
}

type customerPonesImp struct {

}

func (c customerPonesImp)GetCustomerPhoneDetails(info model.CutomerInfo,phone model.Phone,valid bool,number string)(customerPhone model.CustomerPhones){
	if valid{
		customerPhone.State=model.NumberOKState
	}else{
		customerPhone.State=model.NumberNotOKState
	}
	customerPhone.CustomerName =info.Name
	customerPhone.CountryCode = phone.Code
	customerPhone.Country = phone.Country
	customerPhone.PhoneNumber= number

	return customerPhone
}
