package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating User"))
}

func FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find all Users"))
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find User"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
