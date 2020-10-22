package apiController

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	apiServices "api/services"

	// apiModel "api/model"
)

// var TodoList []Todo

type Todo struct {
	Title string
	Content string
	Date string
}

type ApiResponse struct {
	ResultCode    string
	ResultMessage interface{}
}

type Resid struct {
	Id string
}


func AddTodoList(w http.ResponseWriter, r *http.Request){

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024)) //io.LimitReader限制大小
	if err != nil {
		fmt.Println(err)
	}

	var addTodo Todo
	_ = json.Unmarshal(body, &addTodo) //轉為struct 
	// TodoList = append(TodoList, addTodo)
	// TodoList := []Todo{addTodo}
	apiServices.InsertData(body, r)
	defer r.Body.Close()

	response := ApiResponse{"200", addTodo}
	
	
	apiServices.ResponseWithJson(w, http.StatusOK, response) //回傳
}
func GetTodoList(w http.ResponseWriter, r *http.Request){
	// fmt.Println("sjfldf")
	// response := ApiResponse{"200", "dasdw"}
	// apiServices.ResponseWithJson(w, http.StatusOK, response) //回傳
	apiServices.ShowData(w)
}
func UpdateTodoList(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024)) //io.LimitReader限制大小
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()
	var str struct{
		Id int64
		Title string
		Content string
	}
	_ = json.Unmarshal(body, &str)
	
	apiServices.UpdateData(str.Title, str.Content, str.Id)
}
func DeletTodoList(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024)) //io.LimitReader限制大小
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()
	var str struct{Id int64}
	_ = json.Unmarshal(body, &str)
	log.Println((str.Id))

	apiServices.DeletData(str.Id)
}