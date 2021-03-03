package controller

import (
	"net/http"
	"api/services"
)

// todo restful
func AddTodoList(w http.ResponseWriter, r *http.Request){
	body := services.TransPostData(r)
	services.InsertData(body, w)
}

func GetTodoList(w http.ResponseWriter, r *http.Request){
	services.ShowData(w)
}

func UpdateTodoList(w http.ResponseWriter, r *http.Request){
	body := services.TransPostData(r)
	services.UpdateData(body, w)
	
}
func DeletTodoList(w http.ResponseWriter, r *http.Request){
	body := services.TransPostData(r)
	services.DeletData(body, w)
}

// login
func Login(w http.ResponseWriter, r *http.Request){
	body := services.TransPostData(r)
	services.Login(body, w, r)
}

func Logout(w http.ResponseWriter, r *http.Request){
	// body := services.TransPostData(r)
	// services.Logout(body, w, r)
}

func Secret(w http.ResponseWriter, r *http.Request){
	services.Secret(w, r)
}