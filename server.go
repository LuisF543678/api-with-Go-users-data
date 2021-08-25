package main

import (
	"log"
	"fmt"
	//"net/http"
	"encoding/json"
	"github.com/gorilla/mux"

	"./connect"
	"../structures"
	)

/* type User struct{
	id int `json:"id"`
	Username string `json: "username"`
	First_name string `json: "fisrt_name"`
	Last_name string `json: "last_name"`
} */

func main() {
	//fmt.Println("Hola mundo")

	connect.InitializeDatabase()
	defer connect.CloseConnection()
	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", GetUSer).Methods("GET")
	r.HandleFunc("/user/new", GetUSer).Methods("POST")
	r.HandleFunc("/user/aupdate/{id}", GetUSer).Methods("PACHT")
	r.HandleFunc("/user/delete/{id}", DeleteUSer).Methods("DELETE")
	log.Println("El servidor se encuentra en el puerto 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func GetUser(w http.ResponseWriter, r* http.Requests) {
	//w.Write([]byte("Taller REST Gorilla!\n"))

	vars := mux.Vars(r)
	user_id := vars["id"]

	user := connect.GetUSer(user_id)

	status := "success"
	var message string
	user := connect.GetUSer(user_id)
	if(user.Id <= 0 ){
		status = "error"
		message = "User not foud"
	}

	response := structures.Response{ status, user, message}

	//user := User { "nombre apellido", "apellido", "apellido"}
	json.NewEncoder(w).Encode(response)

}

func NewUser(w http.ResponseWriter, r* http.Requests) {
	user := GetUserRequests(r)
	connect.CreateUser(user)
	response := structures.Response{"sucess", connect.CreateUser(user), ""}
	json.NewEncoder(w).Encode(response)
}

func GetUserRequests(r*  http.Requests) structures.User {
	var user structures.USer

	decoder := jsonNewDecoder(r.Body)
	err := decoder.Decode(&user)

	if Err != nil {
		log.Fatal(err)
	}
	returnuser
}

func UpdateUSer(w http.ResponseWriter, r* http.Requests) {
	vars := mux.Vars(r)
	user_id := vars["id"]

	user := GetUserRequests(r)
	response := structures.Response{"sucess", connect.UpdateUSer(user_id, user), ""}
	//connect.UpdateUSer(user_id, user)
	json.NewEncoder(w).Encode(response)
}

func DeleteUSer() {
	vars := mux.Vars(r)
	user_id := vars["id"]

	var user structures.USer
	connect.DeleteUSer(user_id)
	response := structures.Response{"sucess", user, "" }
	json.NewEncoder(w).Encode(response)
}