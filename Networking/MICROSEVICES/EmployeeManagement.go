/*
Q2) Create a simple REST API for managing a list of employees. The API should have the following endpoints:

    GET /employees: Retrieves a list of all employees
    GET /employees/{id}: Retrieves a specific employee by its ID
    POST /employees: Creates a new employee
    PUT /employees/{id}: Updates an existing employee
    DELETE /employees/{id}: Deletes an existing employee
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type EMPLOYEE struct {
	EMP_ID   int
	EMP_NAME string
}

var employee []EMPLOYEE

func main() {

	emp = []EMPLOYEE{
		{EMP_ID: 1, EMP_NAME: "Harry"},
		{EMP_ID: 2, EMP_NAME: "Glory"},
	}

	http.HandleFunc("/employee", handleBooks)
	http.HandleFunc("/emp/", handleBuk)
	http.ListenAndServe(":8082", nil)
}

func handleBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EMPLOYEE :: ")

	switch r.Method {
	case "GET":
		handleGetEmployee(w, r)
	case "POST":
		handlePostEmployee(w, r)
	//case "PUT":
	//case "DELETE":
	default:
		http.Error(w, "INVALID request method", http.StatusMethodNotAllowed)
	}
}

func handleGetEmployee(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(employee)
}

func handlePostEmployee(w http.ResponseWriter, r *http.Request) {
	var em EMPLOYEE
	//read from the request
	json.NewDecoder(r.Body).Decode(&em)
	employee = append(employee, em)
	json.NewEncoder(w).Encode(em) //optional
}

func handleBuk(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle book item api was called")
	//converting string into int
	//ranging whether its present or not
	id, err := strconv.Atoi(r.URL.Path[len("/emp/"):])
	fmt.Println(r.URL.Path[len("/emp/"):])
	path := "/emp/1"
	empid := path[8:]
	fmt.Println(empid)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for i, e := range employee {
		if e.EMP_ID == id {
			switch r.Method {
			case "GET":
				json.NewEncoder(w).Encode(e)
			case "PUT":
				var newEmp EMPLOYEE
				json.NewDecoder(r.Body).Decode(&newEmp)
				newEmp.EMP_ID = id
				employee[i] = newEmp
				json.NewEncoder(w).Encode(newEmp)
			case "DELETE":
				employee = append(employee[:i], employee...)
				w.WriteHeader(http.StatusNoContent)
			default:
				http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			}
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)

}
