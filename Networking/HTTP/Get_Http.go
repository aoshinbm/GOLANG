package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	id   int
	name string
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/personDetails", handlePostRequest)
	http.HandleFunc("/person", userInfo)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

/*
	func handleRoot(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("---------- WELCOME TO USER REGISTRATION ----------"))
	}
*/
func handlePostRequest(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		var person1 Person
		err := json.NewDecoder(request.Body).Decode(&person1)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(writer, "Invalid request: %v", err)
			return
		}
		people = append(people, person1)
		writer.WriteHeader(http.StatusCreated)
		fmt.Println("Item was successfully added to slice.")
	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("Invalid Request-Method Not Found")
	}
}

func userInfo(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		json.NewEncoder(writer).Encode(users)
	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(writer, "Invalid method")
	}
}
