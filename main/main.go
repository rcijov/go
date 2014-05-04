package main

import (
	"fmt"
	"web/main/router"
) 

func main() {

	fmt.Println("http://localhost:1234")
	start := router.Router{}
	start.Listen()
}
