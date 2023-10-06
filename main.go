package main

import (
	r "WebShopping/src/routes"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port 12345")
	r.LoadRoutes()
	//	make server run on port 8080
	// and show print statement in terminal when server is running indicating port number
	http.ListenAndServe(":12345", nil)
	fmt.Println("Server is running on port 12345")
}
