package apiController

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	apiServices "api/services"
)

func AddTodoList(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024)) //io.LimitReader限制大小
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	apiServices.InsertData(body, w)
}

func GetTodoList(w http.ResponseWriter, r *http.Request){
	apiServices.ShowData(w)
}

func UpdateTodoList(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024)) //io.LimitReader限制大小
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	apiServices.UpdateData(body, w)
	
}
func DeletTodoList(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024)) //io.LimitReader限制大小
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	apiServices.DeletData(body, w)
}