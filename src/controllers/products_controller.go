package controllers

import (
	"WebShopping/src/models"
	"fmt"
	"strconv"

	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func IndexHandler(writer http.ResponseWriter, request *http.Request) {
	products := models.GetProducts()
	templates.ExecuteTemplate(writer, "Index", products)
}

func NewHandler(writer http.ResponseWriter, request *http.Request) {
	err := templates.ExecuteTemplate(writer, "New", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func Insert(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Inserting new product...")
	if request.Method != "POST" {
		writer.WriteHeader(400)
	}
	name := request.FormValue("name")
	description := request.FormValue("description")
	price := request.FormValue("price")
	quantity := request.FormValue("quantity")

	if name == "" || description == "" || price == "" || quantity == "" {
		panic("Invalid data!")
		writer.WriteHeader(400)
	}

	floatPrice, priceErr := strconv.ParseFloat(price, 64)
	if priceErr != nil {
		panic("Invalid price!")
		writer.WriteHeader(400)
	}
	intQuantity, quantityErr := strconv.Atoi(quantity)
	if quantityErr != nil {
		panic("Invalid quantity!")
		writer.WriteHeader(400)
	}

	models.CreateProduct(name, description, floatPrice, intQuantity)
	http.Redirect(writer, request, "/", 301)

}

func Delete(writer http.ResponseWriter, request *http.Request) {

	id := request.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		panic("Invalid id!")
		writer.WriteHeader(400)
	}
	models.DeleteProductById(idInt)
	http.Redirect(writer, request, "/", 301)
}

func Edit(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Editing product...")
	id := request.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		panic("Invalid id!")
		writer.WriteHeader(400)
	}
	product := models.GetProductById(idInt)

	fmt.Println(product)
	templates.ExecuteTemplate(writer, "Edit", product)
}

func Update(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Updating product...")
	if request.Method != "POST" {
		writer.WriteHeader(400)
	}
	id := request.FormValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		panic("Invalid id!")
		writer.WriteHeader(400)
	}
	name := request.FormValue("name")
	description := request.FormValue("description")
	price := request.FormValue("price")
	quantity := request.FormValue("quantity")

	if name == "" || description == "" || price == "" || quantity == "" {
		panic("Invalid data!")
		writer.WriteHeader(400)
	}

	floatPrice, priceErr := strconv.ParseFloat(price, 64)
	if priceErr != nil {
		panic("Invalid price!")
		writer.WriteHeader(400)
	}
	intQuantity, quantityErr := strconv.Atoi(quantity)
	if quantityErr != nil {
		panic("Invalid quantity!")
		writer.WriteHeader(400)
	}

	models.EditProductById(idInt, name, description, floatPrice, intQuantity)
	http.Redirect(writer, request, "/", 301)
}
