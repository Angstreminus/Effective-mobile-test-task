package router

import "net/http"

type Router struct {
	Router *http.ServeMux
}

func (r *Router) CreateUser(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) EditUser(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) DeleteUser(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) GetUsers(w http.ResponseWriter, req *http.Request) {
}
