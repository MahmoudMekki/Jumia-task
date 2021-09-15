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

func init() {
	db := database.OpenDB()
	customers := repo.NewCustomersRepo(db)
	phones := repo.NewPhoneRepo()
	customerPhons := repo.NewCustomerPhonesRepo()
	newService = handlers.NewService(customers, phones, customerPhons)
}

func main() {
	newService.Router = mux.NewRouter() //Factor new router
	newService.InitRoutes()
	log.Print("http: --> start listening on host:localhost and port:8500")
	log.Panic(http.ListenAndServe("localhost:8500", newService.Router)) // start the server
}
