package controllers

import (
        "net/http"
	"html/template"
  	"path"	
)

var lp = path.Join("..", "views", "layout.html")
var fp = path.Join("..", "views", "index.html")

type Home struct {}

type Profile struct {
  Name    string
  Email   string
  Phone	  int
}

func (a Home) Home(w http.ResponseWriter) {
       
	profile := Profile{"Alex", "dada@dada.com", 12345}
	
        // ParseFiles
        tmpl,err := template.ParseFiles(lp, fp)
	if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        if err := tmpl.Execute(w, profile); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
        }
}


