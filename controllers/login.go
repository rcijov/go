package controllers

import (
	"net/http"
	"io"
)

type Login struct {
	flag int
}

func (a Login) Login(w http.ResponseWriter) int {
	io.WriteString(w,"Login")
	return a.flag	
}
