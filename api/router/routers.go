package routers

import (
	"net/http"
	apiController "api/controller"
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
	register("POST", "/api/todo/add", apiController.AddTodoList, nil)
	register("GET", "/api/todoList", apiController.GetTodoList, nil)
	register("POST", "/api/todo/delete", apiController.DeletTodoList, nil)
	register("POST", "/api/todo/update", apiController.UpdateTodoList, nil)
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
