package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"httpserver/internal/dto"
	"log"
	"net/http"
	"time"
)

// GetInventory - get all inventory
func GetInventory(w http.ResponseWriter, r *http.Request) {
	url := "http://127.0.0.1:8181/GET/inventory"
	var netClient = http.Client{
		Timeout: time.Second * 10,
	}
	res, err := netClient.Get(url)
	if err != nil {
		log.Println(err)
		return
	}

	var inventory []dto.InventoryStruct
	errDecode := json.NewDecoder(res.Body).Decode(&inventory)
	if errDecode != nil {
		log.Println(errDecode)
		return
	}

	errBodyClose := res.Body.Close()
	if errBodyClose != nil {
		log.Println(errBodyClose)
		return
	}

	tmpl, err := template.ParseFiles("tmpl//getInventory.html")
	if err != nil {
		log.Println("GetInventory: parse err ", err.Error())
		http.Error(w, "Server error", 500)
		return
	}
	err = tmpl.Execute(w, inventory)
	if err != nil {
		log.Println("GetInventory: execute err ", err.Error())
		http.Error(w, "Server error", 500)
		return
	}
}

// CreateInventory - creates inventory
func CreateInventory(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "tmpl//createInventory.html")
	case "POST":
		url := "http://127.0.0.1:8181/POST/create/inventory"
		var netClient = http.Client{
			Timeout: time.Second * 10,
		}
		inventory := dto.InventoryStruct{
			NameOfInventory:        r.FormValue("name"),
			DescriptionOfInventory: r.FormValue("description"),
		}

		invBytes, err := json.Marshal(inventory)
		if err != nil {
			log.Println(err)
			return
		}
		res, err2 := netClient.Post(url, "application/json", bytes.NewBuffer(invBytes))
		if err2 != nil {
			log.Println(err2)
			http.Error(w, "Server error", 500)
			return
		}
		errBodyClose := res.Body.Close()
		if errBodyClose != nil {
			log.Println(errBodyClose)
			http.Error(w, "Server error", 500)
			return
		}
		http.Redirect(w, r, "/GET/inventory", http.StatusMovedPermanently)
	}
}

// DeleteInventory - deletes inventory
func DeleteInventory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hu")
	vars := mux.Vars(r)
	invName := vars["name"]
	url := "http://127.0.0.1:8181/GET/delete/inventory/" + invName
	var netClient = http.Client{
		Timeout: time.Second * 10,
	}

	res, err := netClient.Get(url)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server error", 500)
		return
	}

	errBodyClose := res.Body.Close()
	if errBodyClose != nil {
		log.Println(errBodyClose)
		http.Error(w, "Server error", 500)
		return
	}
	http.Redirect(w, r, "/GET/inventory", http.StatusMovedPermanently)
}

// CreateAssign - create assign
func CreateAssign(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "tmpl//createAssign.html")
	case "POST":
		url := "http://127.0.0.1:8181/POST/create/assign"
		var netClient = http.Client{
			Timeout: time.Second * 10,
		}

		newStruct := struct {
			User      string
			Inventory string
		}{}
		newStruct.User = r.FormValue("user")
		newStruct.Inventory = r.FormValue("inventory")

		invBytes, err := json.Marshal(newStruct)
		if err != nil {
			log.Println(err)
			return
		}
		res, err2 := netClient.Post(url, "application/json", bytes.NewBuffer(invBytes))
		if err2 != nil {
			log.Println(err2)
			http.Error(w, "Server error", 500)
			return
		}
		errBodyClose := res.Body.Close()
		if errBodyClose != nil {
			log.Println(errBodyClose)
			http.Error(w, "Server error", 500)
			return
		}
		http.Redirect(w, r, "/GET/inventory", http.StatusMovedPermanently)
	}
}
