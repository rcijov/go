package router

import (
	"log"
	"web/controllers"
	"net/http"
)

type Router struct {
    flag int
}

// Home Controller
func Home(w http.ResponseWriter, req *http.Request) {
	l := controllers.Home{}
	l.Home(w)
}

// Login Controller
func Login(w http.ResponseWriter, req *http.Request) {
        l := controllers.Login{}
        l.Login(w)
}

// List Controller
func List(w http.ResponseWriter, req *http.Request) {
        l := controllers.List{}
        l.List(w)
}

func (a Router) Listen() int {
	http.HandleFunc("/", Home)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/list", List)
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	return a.flag
}

