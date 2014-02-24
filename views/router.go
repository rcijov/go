package views

import (
	"log"
	"web/controllers"
	"net/http"
	"html/template"
  	"path"
)

type Router struct {
    flag int
}

type Profile struct {
     Name string
     Hobbies []string
}

// Home Controller
func Home(w http.ResponseWriter, req *http.Request) {
	c := controllers.Home{}
	c.Home()
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}	
	lp := path.Join("templates", "layout.html")
        fp := path.Join("templates", "index.html")

        // Note that the layout file must be the first parameter in ParseFiles
        tmpl, err := template.ParseFiles(lp, fp)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }

        if err := tmpl.Execute(w, profile); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
        }

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

// Close Connection
func Close(w http.ResponseWriter, req *http.Request) {
        req.Body.Close()
}


func (a Router) Listen() int {
	http.HandleFunc("/", Home)
	http.HandleFunc("/close", Close)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/list", List)
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	return a.flag
}

