/*
Write a Go program that creates a HTTP server for a e-commerce website,
that listens on
a specific port, and responds to incoming HTTP requests
by fetching product information from a database and returning it to the client in a JSON format.
*/

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

var db *sql.DB

func main() {
	var err error
	//"UserName:Password@tcp(portNumber)/databaseName"
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ecommerce")
	if err != nil {
		fmt.Println("some error")
	}

	http.HandleFunc("/products", productsHandler)

	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}
func productsHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch product from database
	product := Product{
		ID:    1,
		Name:  "Product 1",
		Price: 10,
	}

	// Convert product to JSON
	productJSON, _ := json.Marshal(product)

	// Write JSON to response
	w.Header().Set("Content-Type", "application/json")
	w.Write(productJSON)
}

func listAll() ([]Product, error) {
	var pro []Product
	rows, err := db.Query("SELECT * FROM products; ")
	if err != nil {
		return nil, fmt.Errorf("error in query all contact: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var prod Product
		if err := rows.Scan(&prod.ID, &prod.Name, &prod.Price); err != nil {
			return nil, fmt.Errorf("error in query all contact: %v",
				err)
		}
		pro = append(pro, prod)
	}
	return pro, nil
}
