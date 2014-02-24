package main

import (
	"fmt"
	"web/models"
	"web/views"
)

func main() {
	a := models.Account{}
	a.SetPhone("123")
	fmt.Println(a.GetPhone())
	
	b := views.Router{}
	b.Listen()
}
