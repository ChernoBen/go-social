package controllers

import "net/http"

//Insert an User in DB
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating an User"))
}

//Get all User on DB
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting Users"))
}

//Get one user by ID
func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting an User"))
}

//Update an User by passing his ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating an User"))
}

//Delete an User by passing his ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting an User"))
}
