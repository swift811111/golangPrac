package services

import (
	"net/http"
	"log"
	"api/model"
	"encoding/json"
)

func Login(data []byte,w http.ResponseWriter, r *http.Request){
	usersData := new(model.UsersInfo)
	err := json.Unmarshal(data, usersData) //轉為struct
	if err != nil {
		panic(err.Error())
	}
	ok, result := usersData.Login(w, r)
	if ok {
		ResponseWithJson(w, http.StatusOK, ApiResponse{"200", result})
	}else {
		ResponseWithJson(w, http.StatusOK, ApiResponse{"401", result})
	}
}

func Logout(data []byte,w http.ResponseWriter, r *http.Request){
	// usersData := new(model.UsersInfo)
	// err := json.Unmarshal(data, usersData) //轉為struct
	// if err != nil {
	// 	panic(err.Error())
	// }
	// ok := usersData.LogoutSession(w,r)
	// if ok {
	// 	ResponseWithJson(w, http.StatusOK, ApiResponse{"200", true})
	// }else {
	// 	ResponseWithJson(w, http.StatusOK, ApiResponse{"401", false})
	// }
}
func Secret(w http.ResponseWriter, r *http.Request){
	log.Println("dsf")
	ok := model.Secret(w, r)
	log.Println(ok)
	if ok {
		ResponseWithJson(w, http.StatusOK, ApiResponse{"200", "login"})
	}else {
		ResponseWithJson(w, http.StatusForbidden, ApiResponse{"401", "not login"})
	}
}