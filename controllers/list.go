package controllers

import (
	"net/http"
	"io"
	"web/models"
)

type List struct {
	flag int
}

type Account struct {
    Name string
    Email string
    Phone string
    Password string
    dateCreated string
}

func (a List) List(w http.ResponseWriter) int {
	b := models.Account{}
	io.WriteString(w, b.Account().Name)
	return 0	
}
