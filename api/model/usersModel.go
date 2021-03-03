package model

import (
    _ "github.com/go-sql-driver/mysql"
	"os"
	"golang.org/x/crypto/bcrypt"
	"github.com/gorilla/sessions"
	"net/http"
	"log"
)

type UsersInfo struct {
	Id int64 `json:"id"` 
	Username string `json:"username"`
	Password string `json:"password"`
	Date string `json:"date"`
}

var (
    // key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
    key = []byte(os.Getenv("SESSION_USER_KEY"))
    sessionStore = sessions.NewCookieStore(key)
)

func (e *UsersInfo)Login(w http.ResponseWriter, r *http.Request) (bool, interface{}){
	if e.CheckLogin(){
		e.loginSession(w, r)
		e.Password = ""
		return true, e
	} 
	e.Password = ""
	return false, e
}

func (e *UsersInfo)CheckLogin() bool{
	db := DbConn()
	defer db.Close()
	results, err := db.Query("SELECT id, username, password FROM users WHERE username = ?", e.Username)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
	}
	var tmpusername, tmppassword string
	for results.Next() {
        
        err = results.Scan(&e.Id, &tmpusername, &tmppassword)
        if err != nil {
            panic(err.Error()) 
        }
	}
	return CheckPasswordHash(e.Password, tmppassword)
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func (e *UsersInfo)loginSession(w http.ResponseWriter, r *http.Request){	
	
	session, _ := sessionStore.Get(r, "sessionName")
	session.Values["authenticated"] = true
	session.Save(r, w)
	// log.Println(session.Values["authenticated"])
	// log.Println(session)
	// log.Println(session.Values["cookieName"])
	// session, _ := sessionStore.Get(r, "sessionName")
	// session.Values["sessionName"] = e.Username
	// err := session.Save(r, w)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}

func Secret(w http.ResponseWriter, r *http.Request) bool{
    session, _ := sessionStore.Get(r, "sessionName")

    // Check if user is authenticated
    // if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
    //     // http.Error(w, "Forbidden", http.StatusForbidden)
    //     return false
	// }
	log.Println(session.Values["authenticated"])
	if session.Values["authenticated"] != true {
		return false
	}
	return true
    // Print secret message
    // fmt.Fprintln(w, "The cake is a lie!")
}

// func (e *UsersInfo)LogoutSession(w http.ResponseWriter, r *http.Request) bool{
// 	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_USER_KEY")))
// 	session, _ := store.Get(r, e.Username)
// 	delete(session.Values, e.Username)
// 	// err := session.Save(r, w)
// 	// if err != nil {
// 	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 	// 	return false
// 	// }
// 	return true
// }