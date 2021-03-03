package routers

import (
	"net/http"
	"api/controller"
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
)

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routers []Route

func init() {
	godotenv.Load()
	register("POST", "/api/todo/add", controller.AddTodoList, nil)
	register("GET", "/api/todo/showList", controller.GetTodoList, nil)
	register("POST", "/api/todo/delete", controller.DeletTodoList, nil)
	register("POST", "/api/todo/update", controller.UpdateTodoList, nil)
	register("POST", "/api/todo/login", controller.Login, nil)
	register("POST", "/api/todo/logout", controller.Logout, nil)
	register("POST", "/api/todo/secret", controller.Secret, nil)
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	for _, route := range routers {
		r.Methods(route.Method).
			Path(route.Pattern).
			Handler(route.Handler)
		if route.Middleware != nil {
			r.Use(route.Middleware)
		}
	}
	return r
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routers = append(routers, Route{method, pattern, handler, middleware})
}
