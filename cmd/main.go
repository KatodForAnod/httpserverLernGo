package main

import (
	"github.com/gorilla/mux"
	"httpserver/internal/handler"
	"log"
	"net/http"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/GET/users", handler.GetUsers)
	mux.HandleFunc("/POST/create/user", handler.CreateUser)
	mux.HandleFunc("/GET/delete/user/{name:[a-zA-z]+}", handler.DeleteUser)

	mux.HandleFunc("/GET/inventory", handler.GetInventory)
	mux.HandleFunc("/POST/create/inventory", handler.CreateInventory)
	mux.HandleFunc("/GET/delete/inventory/{name:[a-zA-z]+}", handler.DeleteInventory)

	mux.HandleFunc("/POST/create/assign", handler.CreateAssign)
	http.Handle("/", mux)

	log.Println("Запуск веб-сервера на http://127.0.0.1:8080/POST/create/assign")
	errServe := http.ListenAndServe(":8080", nil)
	log.Fatal(errServe)
}
