package apiController

import (
	"net/http"
	apiServices "api/services"
)

func AddTodoList(w http.ResponseWriter, r *http.Request){
	body := apiServices.TransPostData(r)
	apiServices.InsertData(body, w)
}

func GetTodoList(w http.ResponseWriter, r *http.Request){
	apiServices.ShowData(w)
}

func UpdateTodoList(w http.ResponseWriter, r *http.Request){
	body := apiServices.TransPostData(r)
	apiServices.UpdateData(body, w)
	
}
func DeletTodoList(w http.ResponseWriter, r *http.Request){
	body := apiServices.TransPostData(r)
	apiServices.DeletData(body, w)
}