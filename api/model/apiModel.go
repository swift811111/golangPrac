package apiModel

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
    "os"
)

type ShowTodo struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Date string `json:"date"`
}

func DbConn() (db *sql.DB) {
    db, err := sql.Open(os.Getenv("dbDriver"), os.Getenv("dbUser")+":"+os.Getenv("dbPass")+"@tcp(database-1.c7vvkbv4olla.ap-northeast-1.rds.amazonaws.com:3306)/"+os.Getenv("dbName")+"?charset=utf8")
    if err != nil {
        panic(err.Error())
	}
    return db
}

func (e *ShowTodo)InserData(){
	db := DbConn()
	defer db.Close()

	title := e.Title
    content := e.Content
    insForm, err := db.Prepare("INSERT INTO todoList(title, content, date) VALUES(?,?,?)")
    if err != nil {
        panic(err.Error())
    }
    insForm.Exec(title, content, "NOW()")
}

func (e *ShowTodo)GetData() []ShowTodo{
	db := DbConn()
    defer db.Close()

	results, err := db.Query("SELECT id, title, content, date FROM todoList")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
	}

	var resultArry []ShowTodo
	for results.Next() {
        // for each row, scan the result into our tag composite object
        err = results.Scan(&e.Id, &e.Title, &e.Content, &e.Date)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }

		resultArry = append(resultArry, *e)
    }
    return resultArry
}

func (e *ShowTodo)UpdateData(){
    db := DbConn()
    defer db.Close()
    _, err := db.Query("UPDATE todoList set title=?, content=? where id=?",e.Title, e.Content, e.Id)
    if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
    }
}

func (e *ShowTodo)DeletData(){
    db := DbConn()
    defer db.Close()

    result,err := db.Exec("delete from todoList where id=?",e.Id)
    if err != nil{
        fmt.Printf("delete failed,err:%v\n",err)
        return
    }
    
    fmt.Println("delete data successd:", result)
	fmt.Println("delete item id:", e.Id)
}