package controllers

import (
	"net/http"
	"io"
)

type List struct {
	flag int
}

func (a List) List(w http.ResponseWriter) int {
	io.WriteString(w,"List")
	return a.flag	
}
