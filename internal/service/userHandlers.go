package service

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"httpserver/internal/dto"
	"log"
	"net/http"
	"time"
)

// GetUsers - return users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	url := "http://127.0.0.1:8181/users"
	var netClient = http.Client{
		Timeout: time.Second * 10,
	}
	res, err := netClient.Get(url)
	if err != nil {
		log.Println(err)
		return
	}

	var users []dto.UserStruct
	errDecode := json.NewDecoder(res.Body).Decode(&users)
	if errDecode != nil {
		log.Println(errDecode)
		return
	}

	errBodyClose := res.Body.Close()
	if errBodyClose != nil {
		log.Println(errBodyClose)
		return
	}

	tmpl, err := template.ParseFiles("tmpl//getUsers.html")
	if err != nil {
		log.Println("GetUsers: parse err ", err.Error())
		http.Error(w, "Server error", 500)
		return
	}
	err = tmpl.Execute(w, users)
	if err != nil {
		log.Println("GetUsers: execute err ", err.Error())
		http.Error(w, "Server error", 500)
		return
	}
}

// CreateUser - create user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "tmpl//createUser.html")
	case "POST":
		url := "http://127.0.0.1:8181/user"
		var netClient = http.Client{
			Timeout: time.Second * 10,
		}
		user := dto.UserStruct{
			UserName: r.FormValue("UserName"),
		}
		productBytes, err := json.Marshal(user)
		if err != nil {
			log.Println(err)
			return
		}
		res, err2 := netClient.Post(url, "application/json", bytes.NewBuffer(productBytes))
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
		http.Redirect(w, r, "/users", http.StatusMovedPermanently)
	}
}

// DeleteUser - deletes users
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["name"]
	url := "http://127.0.0.1:8181/user/" + userName
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
	http.Redirect(w, r, "/users", http.StatusMovedPermanently)
}
