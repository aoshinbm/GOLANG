/*
Q1)   Create a simple REST API for managing a list of books. The API should have the following endpoints:

	GET /books: Retrieves a list of all books
	GET /books/{id}: Retrieves a specific book by its ID
	POST /books: Creates a new book
	PUT /books/{id}: Updates an existing book
	DELETE /books/{id}: Deletes an existing book
*/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type BOOK struct {
	BOOK_ID   int
	BOOK_NAME string
}

var books []BOOK

func main() {

	books = []BOOK{
		{BOOK_ID: 1, BOOK_NAME: "Harry Potter"},
		{BOOK_ID: 2, BOOK_NAME: "House Of Dragons"},
	}

	http.HandleFunc("/books", handleBooks)
	http.HandleFunc("/books/", handleBuk)
	http.ListenAndServe(":8081", nil)
}

func handleBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("BOOKS ")

	switch r.Method {
	case "GET":
		handleGetBooks(w, r)
	case "POST":
		handleBuk(w, r)
	//case "PUT":
	//case "DELETE":
	default:
		http.Error(w, "INVALID request method", http.StatusMethodNotAllowed)
	}
}

func handleGetBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func handlePostStock(w http.ResponseWriter, r *http.Request) {
	var bk BOOK
	//read from the request
	json.NewDecoder(r.Body).Decode(&bk)
	books = append(books, bk)
	json.NewEncoder(w).Encode(bk) //optional
}

func handleBuk(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle book item api was called")
	//converting string into int
	//ranging whether its present or not
	id, err := strconv.Atoi(r.URL.Path[len("/books/"):])
	fmt.Println(r.URL.Path[len("/books/"):])
	path := "/books/12"
	bookid := path[8:]
	fmt.Println(bookid)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for i, bookk := range books {
		if bookk.BOOK_ID == id {
			switch r.Method {
			case "GET":
				json.NewEncoder(w).Encode(bookk)
			case "PUT":
				var newBook BOOK
				json.NewDecoder(r.Body).Decode(&newBook)
				newBook.BOOK_ID = id
				books[i] = newBook
				json.NewEncoder(w).Encode(newBook)
			case "DELETE":
				books = append(books[:i], books...)
				w.WriteHeader(http.StatusNoContent)
			default:
				http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			}
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)

}
