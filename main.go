package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	_ "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"lecture8/database"
	"net/http"
)
func main(){
	router := mux.NewRouter().StrictSlash(true)
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*",})

	router.Methods(http.MethodGet).Path("/people/{id}").HandlerFunc(database.GetPeopleId)
	router.Methods(http.MethodGet).Path("/people").HandlerFunc(database.GetAllPeople)
	router.Methods(http.MethodPost).Path("/people").HandlerFunc(database.CreatePeople)
	router.Methods(http.MethodPut).Path("/people/{id}").HandlerFunc(database.UpdatePeople)
	router.Methods(http.MethodDelete).Path("/people/{id}").HandlerFunc(database.DeletePeople)

	err := http.ListenAndServe(":9000", handlers.CORS(headers, methods, origins)(router))
	if err != nil {
		panic(err)
	}
}