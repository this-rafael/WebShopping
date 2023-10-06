package models

import (
	"WebShopping/src/config"
	"database/sql"
	"fmt"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func getProductsQuery() string {
	return "SELECT id, name, description, price, quantity FROM products"
}

func queryProducts(db *sql.DB) *sql.Rows {
	queryProducts, err := db.Query(getProductsQuery())
	if err != nil {
		panic(err)
	}
	return queryProducts
}

func mapSqlRowsToProduct(foundedRows *sql.Rows) []Product {
	products := []Product{}
	tempProduct := Product{}

	for foundedRows.Next() {
		var id int
		var name string
		var description string
		var price float64
		var quantity int

		err := foundedRows.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err)
		}
		tempProduct.Id = id
		tempProduct.Name = name
		tempProduct.Description = description
		tempProduct.Price = price
		tempProduct.Quantity = quantity
		products = append(products, tempProduct)
	}
	return products
}

func GetProducts() []Product {
	db := config.DatabaseConnect()
	foundedRows := queryProducts(db)
	products := mapSqlRowsToProduct(foundedRows)

	defer db.Close()

	return products

}

func getProductByIdQuery() string {
	return "SELECT id, name, description, price, quantity FROM products WHERE id=$1 LIMIT 1"
}

func queryProductById(db *sql.DB, id int) *sql.Rows {
	queryProduct, err := db.Query(getProductByIdQuery(), id)
	if err != nil {
		panic(err)
	}
	return queryProduct
}

func GetProductById(id int) Product {
	db := config.DatabaseConnect()
	foundedRow := queryProductById(db, id)
	fmt.Println(foundedRow)
	product := mapSqlRowsToProduct(foundedRow)[0]
	defer db.Close()
	return product
}

func insertProductQuery() string {
	return "INSERT INTO products(name, description, price, quantity) VALUES($1, $2, $3, $4)"
}

func CreateProduct(name string, description string, price float64, quantity int) {
	db := config.DatabaseConnect()
	insertProduct, err := db.Prepare(insertProductQuery())
	if err != nil {
		panic(err)
	}
	_, err2 := insertProduct.Exec(name, description, price, quantity)
	if err2 != nil {
		panic(err2)
	}

	defer db.Close()
}

func deleteProductByIdQuery() string {
	return "DELETE FROM products WHERE id=$1"
}

func DeleteProductById(id int) {
	db := config.DatabaseConnect()
	deleteProduct, err := db.Prepare(deleteProductByIdQuery())
	if err != nil {
		panic(err)
	}
	result, err2 := deleteProduct.Exec(id)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println(result)
	defer db.Close()
}

func editProductByIdQuery() string {
	return "UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id=$5"
}

func EditProductById(id int, name string, description string, price float64, quantity int) {
	db := config.DatabaseConnect()
	editProduct, err := db.Prepare(editProductByIdQuery())
	if err != nil {
		panic(err)
	}
	result, err2 := editProduct.Exec(name, description, price, quantity, id)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println(result)
	defer db.Close()
}
