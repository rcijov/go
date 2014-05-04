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
    Type string
    Date string
}

func (a List) List(w http.ResponseWriter) int {
	b := models.Account{}
	s := "Name: " + b.Account().Name
	s += "\nPhone: " + b.Account().Phone
	s += "\nPassword: " + b.Account().Password
	s += "\nEmail: " + b.Account().Email
	s += "\nType: " + b.Account().Type
	s += "\nDate: " + b.Account().Date.String()
	io.WriteString(w, s)
	return 0	
}
