package handlers

import (
	"github.com/gorilla/mux"
	"jumia-task/pkg/repo"
	"net/http"
)

type Service struct {
	CustomersRepo  repo.CustomersRepo
	PhoneRepo	   repo.PhoneRepo
	CustomerPhonesRepo	repo.CustomerPhonesRepo
	Router 				*mux.Router
}

func NewService(customer repo.CustomersRepo,phone repo.PhoneRepo,customerPhone repo.CustomerPhonesRepo) *Service{
	return &Service{
		CustomersRepo: customer,
		PhoneRepo: phone,
		CustomerPhonesRepo: customerPhone,
	}
}

func (s *Service)InitRoutes(){
	s.Router.HandleFunc("/customers",s.GetAllCustomerPhones).Methods(http.MethodGet)
	s.Router.HandleFunc("/countries",s.GetAllAvailableCountriesForPagination).Methods(http.MethodGet)
	s.Router.HandleFunc("/customers-pagination",s.GetCustomerPhonesByPagination).Methods(http.MethodPost)
}