package router

import (
	"log"
	"net/http"

	librarycontroller "github.com/Sutrisno-14/perpustakaan-restful-api/controller/libraryController"
	"github.com/gorilla/mux"
)

func Rooter() {
	r :=mux.NewRouter()
	port := ":8081"

	//book
	r.HandleFunc("/api/books", librarycontroller.CreateBook).Methods("POST")
	r.HandleFunc("/api/books", librarycontroller.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", librarycontroller.GetBookById).Methods("GET")
	r.HandleFunc("/api/books/{id}", librarycontroller.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books", librarycontroller.DeleteBook).Methods("DELETE")

	//employee
	r.HandleFunc("/api/employees", librarycontroller.CreateEmployee).Methods("POST")
	r.HandleFunc("/api/employees", librarycontroller.GetEmployees).Methods("GET")
	r.HandleFunc("/api/employees/{id}", librarycontroller.GetEmployeeById).Methods("GET")
	r.HandleFunc("/api/employees/{id}", librarycontroller.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/api/employees", librarycontroller.DeleteEmployee).Methods("DELETE")

	//visitor
	r.HandleFunc("/api/visitors", librarycontroller.CreateVisitor).Methods("POST")
	r.HandleFunc("/api/visitors", librarycontroller.GetVisitors).Methods("GET")
	r.HandleFunc("/api/visitors/{id}", librarycontroller.GetVisitorById).Methods("GET")
	r.HandleFunc("/api/visitors/{id}", librarycontroller.UpdateVisitor).Methods("PUT")
	r.HandleFunc("/api/visitors", librarycontroller.DeleteVisitor).Methods("DELETE")

	log.Printf("Server running port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}