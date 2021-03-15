package main

import (
	"github.com/gorilla/mux"
	"httpserver/internal/service"
	"log"
	"net/http"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/users", service.GetUsers)
	mux.HandleFunc("/user", service.CreateUser)
	mux.HandleFunc("/user/{name:[a-zA-z]+}", service.DeleteUser)

	mux.HandleFunc("/inventories", service.GetInventory)
	mux.HandleFunc("/inventory", service.CreateInventory)
	mux.HandleFunc("/inventory/{name:[a-zA-z]+}", service.DeleteInventory)

	mux.HandleFunc("/assign", service.CreateAssign)
	http.Handle("/", mux)

	log.Println("Запуск веб-сервера на http://127.0.0.1:8080/users")
	errServe := http.ListenAndServe(":8080", nil)
	log.Fatal(errServe)
}
