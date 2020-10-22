package apiServices

import (
	// "log"
	"encoding/json"
	"net/http"
	apiModel "api/model"
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
func InsertData(addTodo []byte, w http.ResponseWriter){
	todoData := new(apiModel.ShowTodo)
	_ = json.Unmarshal(addTodo, todoData) //轉為struct

	todoData.InserData()
	response := ApiResponse{"200", todoData}
	ResponseWithJson(w, http.StatusOK, response)
}

func ShowData(w http.ResponseWriter){
	todoData := new(apiModel.ShowTodo)
	item := todoData.GetData()
	response := ApiResponse{"200", item}
	ResponseWithJson(w, http.StatusOK, response)
}

func UpdateData(data []byte, w http.ResponseWriter){
	todoData := new(apiModel.ShowTodo)
	_ = json.Unmarshal(data, todoData) //轉為struct

	todoData.UpdateData()
	response := ApiResponse{"200", todoData}
	ResponseWithJson(w, http.StatusOK, response)
}

func DeletData(data []byte, w http.ResponseWriter)  {
	todoData := new(apiModel.ShowTodo)
	_ = json.Unmarshal(data, todoData) //轉為struct
	response := ApiResponse{"200", todoData}
	ResponseWithJson(w, http.StatusOK, response)
	todoData.DeletData()
}
