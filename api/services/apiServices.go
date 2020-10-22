package apiServices

import (
	// "log"
	"encoding/json"
	"net/http"
	"fmt"
	apiModel "api/model"
)

type Todo struct {
	Title string
	Content string
	Date string
}

type ShowTodo struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Date string `json:"date"`
}

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func InsertData(addTodo []byte, r *http.Request) {
	db := apiModel.DbConn()

	var addTodoList Todo
	_ = json.Unmarshal(addTodo, &addTodoList) //轉為struct

    if r.Method == "POST" {
		title := addTodoList.Title
		content := addTodoList.Content
        insForm, err := db.Prepare("INSERT INTO todoList(title, content, date) VALUES(?,?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(title, content, "NOW()")
        // log.Println("INSERT: Name: " + name + " | City: " + city)
    }
	defer db.Close()
}

func ShowData(w http.ResponseWriter){
	db := apiModel.DbConn()

	results, err := db.Query("SELECT id, title, content, date FROM todoList")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
	}

	var rearry []ShowTodo
	for results.Next() {
        var showTodo ShowTodo
        // for each row, scan the result into our tag composite object
        err = results.Scan(&showTodo.Id, &showTodo.Title, &showTodo.Content, &showTodo.Date)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
				// and then print out the tag's Name attribute
		rearry = append(rearry, showTodo)
		//log.Printf(showTodo.Title)
		
	}
	ResponseWithJson(w, http.StatusOK, rearry)
	defer db.Close()
}

func UpdateData(title string, content string, id int64)  {
	db := apiModel.DbConn()
	defer db.Close()

	_, err := db.Query("UPDATE todoList set title=?, content=? where id=?",title, content, id)
    if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Println(" update success", id)
}

func DeletData(id int64)  {
	db := apiModel.DbConn()
	defer db.Close()

	result,err := db.Exec("delete from todoList where id=?",id)
    if err != nil{
        fmt.Printf("delete failed,err:%v\n",err)
        return
    }
	fmt.Println("delete data successd:", result)
	fmt.Println("delete item id:", id)
}
