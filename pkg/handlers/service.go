package handlers

import "jumia-task/pkg/repo"

type Service struct {
	CustomersRepo  repo.CustomersRepo
	PhoneRepo	   repo.PhoneRepo
	CustomerPhonesRepo	repo.CustomerPhonesRepo
}

func NewService(customer repo.CustomersRepo,phone repo.PhoneRepo,customerPhone repo.CustomerPhonesRepo) *Service{
	return &Service{
		CustomersRepo: customer,
		PhoneRepo: phone,
		CustomerPhonesRepo: customerPhone,
	}
}