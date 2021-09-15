package handlers

import (
	"github.com/gorilla/mux"
	response "jumia-task/kit"
	model "jumia-task/pkg/models"
	"net/http"
	"strconv"
	"strings"
)

func (s *Service) GetAllCustomerPhones(w http.ResponseWriter, req *http.Request) {
	var customerPhones []model.CustomerPhones
	customers, err := s.CustomersRepo.GetCustomers()
	if err != nil {
		response.RespondError(w,http.StatusInternalServerError,err.Error())
	}

	for _, v := range customers {
		countryKey, phoneNumber := s.PhoneRepo.GetCountryKeyAndPhoneNumber(v.Phone)
		countryPhone, existed := s.PhoneRepo.GetCountryPhoneDetails(countryKey)
		if !existed {
			response.RespondError(w,http.StatusNotFound,"Not suppoerted country")
		}
		valid := s.PhoneRepo.ValidateNumber(v.Phone, countryPhone.Regex)
		customerPhone := s.CustomerPhonesRepo.GetCustomerPhoneDetails(v, countryPhone, valid, phoneNumber)
		customerPhones = append(customerPhones, customerPhone)
	}
	response.RespondJSON(w,http.StatusOK,customerPhones)
}

func (s *Service) GetAllAvailableCountriesForPagination(w http.ResponseWriter, req *http.Request) {
	countries := s.PhoneRepo.GetAvailablePhones()
	response.RespondJSON(w,http.StatusOK,countries)
}

func (s *Service) GetCustomerPhonesByPagination(w http.ResponseWriter, req *http.Request) {
	var customerPhones []model.CustomerPhones
	vars := mux.Vars(req)
	limit, err := strconv.Atoi(vars["limit"])
	if err != nil {
		response.RespondError(w,http.StatusBadRequest,err.Error())
	}
	page, err := strconv.Atoi(vars["page"])
	if err != nil {
		response.RespondError(w,http.StatusBadRequest,err.Error())
	}
	var pagination = model.Pagination{
		FilterBy: vars["filter_by"],
		Limit:    uint32(limit),
		Page:     uint32(page),
	}

	switch strings.ToLower(pagination.FilterBy) {
	case strings.ToLower(model.NumberOKState):
		customers, err := s.CustomersRepo.GetCustomersPagination(pagination)
		if err != nil {
			response.RespondError(w,http.StatusInternalServerError,err.Error())
		}
		for _, v := range customers {
			countryKey, phoneNumber := s.PhoneRepo.GetCountryKeyAndPhoneNumber(v.Phone)
			countryPhone, existed := s.PhoneRepo.GetCountryPhoneDetails(countryKey)
			if !existed {
				response.RespondError(w,http.StatusNotFound,"Not suppoerted country")
			}
			valid := s.PhoneRepo.ValidateNumber(v.Phone, countryPhone.Regex)
			if !valid {
				continue
			}
			customerPhone := s.CustomerPhonesRepo.GetCustomerPhoneDetails(v, countryPhone, valid, phoneNumber)
			customerPhones = append(customerPhones, customerPhone)
		}
		response.RespondJSON(w,http.StatusOK,customerPhones)

	case strings.ToLower(model.NumberNotOKState):
		customers, err := s.CustomersRepo.GetCustomersPagination(pagination)
		if err != nil {
			response.RespondError(w,http.StatusInternalServerError,err.Error())
		}

		for _, v := range customers {
			countryKey, phoneNumber := s.PhoneRepo.GetCountryKeyAndPhoneNumber(v.Phone)
			countryPhone, existed := s.PhoneRepo.GetCountryPhoneDetails(countryKey)
			if !existed {
				response.RespondError(w,http.StatusNotFound,"Not suppoerted country")
			}
			valid := s.PhoneRepo.ValidateNumber(v.Phone, countryPhone.Regex)
			if valid {
				continue
			}
			customerPhone := s.CustomerPhonesRepo.GetCustomerPhoneDetails(v, countryPhone, valid, phoneNumber)
			customerPhones = append(customerPhones, customerPhone)
		}
		response.RespondJSON(w,http.StatusOK,customerPhones)

	default:
		customers, err := s.CustomersRepo.GetCustomersByCountry(pagination)
		if err != nil {
			response.RespondError(w,http.StatusInternalServerError,err.Error())
		}

		for _, v := range customers {
			countryKey, phoneNumber := s.PhoneRepo.GetCountryKeyAndPhoneNumber(v.Phone)
			countryPhone, existed := s.PhoneRepo.GetCountryPhoneDetails(countryKey)
			if !existed {
				response.RespondError(w,http.StatusNotFound,"Not suppoerted country")
			}
			valid := s.PhoneRepo.ValidateNumber(v.Phone, countryPhone.Regex)
			customerPhone := s.CustomerPhonesRepo.GetCustomerPhoneDetails(v, countryPhone, valid, phoneNumber)
			customerPhones = append(customerPhones, customerPhone)
		}
		response.RespondJSON(w,http.StatusOK,customerPhones)
	}
}
