package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	model "jumia-task/pkg/models"
	"log"
	"net/http"
	"strconv"
	"strings"
)



func (s *Service)GetAllCustomerPhones(w http.ResponseWriter,req *http.Request){
	var customerPhones []model.CustomerPhones
	customers,err := s.CustomersRepo.GetCustomers()
	if err !=nil {
		log.Panic("implement the db connection")
	}

	for _,v:= range customers{
		countryKey,phoneNumber := s.PhoneRepo.GetCountryKeyAndPhoneNumber(v.Phone)
		countryPhone,existed := s.PhoneRepo.GetCountryPhoneDetails(countryKey)
		if !existed{
			log.Panic("implement not available country err")
		}
		valid := s.PhoneRepo.ValidateNumber(v.Phone,countryPhone.Regex)
		customerPhone:=s.CustomerPhonesRepo.GetCustomerPhoneDetails(v,countryPhone,valid,phoneNumber)
		customerPhones = append(customerPhones,customerPhone)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customerPhones)
}




func(s *Service)GetAllAvailableCountriesForPagination(w http.ResponseWriter,req *http.Request){
	countries := s.PhoneRepo.GetAvailablePhones()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(countries)
}

func(s *Service)GetCustomerPhonesByPagination(w http.ResponseWriter,req *http.Request){
	var customerPhones []model.CustomerPhones
	vars := mux.Vars(req)
	limit,err := strconv.Atoi(vars["limit"])
	if err !=nil{
		log.Panic("handle not qualified req arguments")
	}
	page,err := strconv.Atoi(vars["page"])
	if err !=nil{
		log.Panic("handle not qualified req arguments")
	}
	var pagination = model.Pagination{
		FilterBy: vars["filter_by"],
		Limit:    uint32(limit),
		Page:     uint32(page),
	}
	switch strings.ToLower(pagination.FilterBy) {
	case strings.ToLower(model.NumberOKState):
		customers,err := s.CustomersRepo.GetCustomersPagination(pagination)
		if err !=nil {
			log.Panic("implement the db connection")
		}

		for _,v:= range customers{
			countryKey,phoneNumber := s.PhoneRepo.GetCountryKeyAndPhoneNumber(v.Phone)
			countryPhone,existed := s.PhoneRepo.GetCountryPhoneDetails(countryKey)
			if !existed{
				log.Panic("implement not available country err")
			}
			valid := s.PhoneRepo.ValidateNumber(v.Phone,countryPhone.Regex)
			if !valid{continue}
			customerPhone:=s.CustomerPhonesRepo.GetCustomerPhoneDetails(v,countryPhone,valid,phoneNumber)
			customerPhones = append(customerPhones,customerPhone)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customerPhones)

	case strings.ToLower(model.NumberNotOKState):
		customers,err := s.CustomersRepo.GetCustomersPagination(pagination)
		if err !=nil {
			log.Panic("implement the db connection")
		}

		for _,v:= range customers{
			countryKey,phoneNumber := s.PhoneRepo.GetCountryKeyAndPhoneNumber(v.Phone)
			countryPhone,existed := s.PhoneRepo.GetCountryPhoneDetails(countryKey)
			if !existed{
				log.Panic("implement not available country err")
			}
			valid := s.PhoneRepo.ValidateNumber(v.Phone,countryPhone.Regex)
			if valid{continue}
			customerPhone:=s.CustomerPhonesRepo.GetCustomerPhoneDetails(v,countryPhone,valid,phoneNumber)
			customerPhones = append(customerPhones,customerPhone)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customerPhones)

	default:
		customers,err := s.CustomersRepo.GetCustomersByCountry(pagination)
		if err !=nil {
			log.Panic("implement the db connection")
		}

		for _,v:= range customers{
			countryKey,phoneNumber := s.PhoneRepo.GetCountryKeyAndPhoneNumber(v.Phone)
			countryPhone,existed := s.PhoneRepo.GetCountryPhoneDetails(countryKey)
			if !existed{
				log.Panic("implement not available country err")
			}
			valid := s.PhoneRepo.ValidateNumber(v.Phone,countryPhone.Regex)
			customerPhone:=s.CustomerPhonesRepo.GetCustomerPhoneDetails(v,countryPhone,valid,phoneNumber)
			customerPhones = append(customerPhones,customerPhone)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customerPhones)

	}
}