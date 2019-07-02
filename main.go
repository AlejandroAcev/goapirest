package main

import (
	"goapi/controller"
	//"goapi/models"
	//"log"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)


func main() {
	fmt.Println("Iniciando api")
	r := mux.NewRouter()
	r.HandleFunc("/hola", controller.HolaHandler).Methods("GET")
	r.HandleFunc("/register", controller.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", controller.LoginHandler).Methods("POST")
	r.HandleFunc("/profile", controller.ProfileHandler).Methods("GET")
	fmt.Println("Iniciadas rutas")

	server := &http.Server{
		Addr:		":8080",
		Handler:	r,
	}

	server.ListenAndServe()

}