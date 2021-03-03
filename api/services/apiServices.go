package services

import (
	// "log"
	"encoding/json"
	"net/http"
	"api/model"
	"io"
	"io/ioutil"
	"fmt"
)

type ApiResponse struct {
	ResultCode    string
	ResultMessage interface{}
}

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func TransPostData(r *http.Request) []byte{
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024)) //io.LimitReader限制大小
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()
	return body
}

func InsertData(addTodo []byte, w http.ResponseWriter){
	todoData := new(model.ShowTodo)
	_ = json.Unmarshal(addTodo, todoData) //轉為struct

	todoData.InserData()
	response := ApiResponse{"200", todoData}
	ResponseWithJson(w, http.StatusOK, response)
}

func ShowData(w http.ResponseWriter){
	todoData := new(model.ShowTodo)
	item := todoData.GetData()
	response := ApiResponse{"200", item}
	ResponseWithJson(w, http.StatusOK, response)
}

func UpdateData(data []byte, w http.ResponseWriter){
	todoData := new(model.ShowTodo)
	_ = json.Unmarshal(data, todoData) //轉為struct
	
	todoData.UpdateData()
	response := ApiResponse{"200", todoData}
	ResponseWithJson(w, http.StatusOK, response)
}

func DeletData(data []byte, w http.ResponseWriter)  {
	todoData := new(model.ShowTodo)
	_ = json.Unmarshal(data, todoData) //轉為struct
	response := ApiResponse{"200", todoData}
	ResponseWithJson(w, http.StatusOK, response)
	todoData.DeletData()
}
