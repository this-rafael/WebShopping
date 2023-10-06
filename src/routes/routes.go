package routes

import (
	"WebShopping/src/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.IndexHandler)
	http.HandleFunc("/new", controllers.NewHandler)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
}
