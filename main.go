package main

import (
	"github.com/gorilla/mux"
	"jumia-task/database"
	"jumia-task/pkg/handlers"
	"jumia-task/pkg/repo"
	"log"
	"net/http"
)

var newService *handlers.Service

func init(){
	db := database.OpenDB()
	customers := repo.NewCustomersRepo(db)
	phones := repo.NewPhoneRepo()
	customerPhons :=repo.NewCustomerPhonesRepo()
	newService = handlers.NewService(customers,phones,customerPhons)
}


func main(){
	r := mux.NewRouter()
	r.HandleFunc("/customers",newService.GetAllCustomerPhones).Methods(http.MethodGet)
	r.HandleFunc("/countries",newService.GetAllAvailableCountriesForPagination).Methods(http.MethodGet)
	r.HandleFunc("/customers-pagination",newService.GetCustomerPhonesByPagination).Methods(http.MethodPost)
	log.Panic(http.ListenAndServe("localhost:8500",r))
}