package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Struct creation for stock
// format of data
type Stock struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

// slice of stock
// storage of data
var stocks []Stock

// main method addition
func main() {

	// Ready made data
	stocks = []Stock{
		{ProductID: 1, Quantity: 10},
		{ProductID: 2, Quantity: 20},
	}

	//2 parameters:- url path & function that will handle request
	http.HandleFunc("/stocks", handleStocks) //  endpoint 1
	http.HandleFunc("/stocks/", handleStock) //  another endpoint  2
	//method from http method
	//listening to http calls n responding to request
	http.ListenAndServe(":8080", nil) // listen and serve
}

// handleStocks
// 2 parameters :- response is interface, request is struct
func handleStocks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle stocks multi item api was called")
	//r.method returns which method is the request calling
	//
	switch r.Method {
	case "GET":
		handleGetStocks(w, r)
	case "POST":
		handlePostStock(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// handle get stocks will return the all stocks
func handleGetStocks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(stocks)
}

// handle Post Stock
func handlePostStock(w http.ResponseWriter, r *http.Request) {
	var stock Stock
	//read from the request
	json.NewDecoder(r.Body).Decode(&stock)
	stocks = append(stocks, stock)
	json.NewEncoder(w).Encode(stock) //optional
}

// another endpoint 2 method
// it will handle all the
// request for particular object

func handleStock(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle stock single item api was called")
	//converting string into int
	//ranging whether its present or not
	id, err := strconv.Atoi(r.URL.Path[len("/stocks/"):])
	fmt.Println(r.URL.Path[len("/stocks/"):])
	path := "/stocks/123"
	stockID := path[8:]
	fmt.Println(stockID)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for i, stock := range stocks {
		if stock.ProductID == id {
			switch r.Method {
			case "GET":
				json.NewEncoder(w).Encode(stock)
			case "PUT":
				var newStock Stock
				json.NewDecoder(r.Body).Decode(&newStock)
				newStock.ProductID = id
				stocks[i] = newStock
				json.NewEncoder(w).Encode(newStock)
			case "DELETE":
				stocks = append(stocks[:i], stocks[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
			default:
				http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			}
			return
		}
	}

	http.Error(w, "Stock not found", http.StatusNotFound)
}
