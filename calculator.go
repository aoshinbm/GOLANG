package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	Operation string  `json:"Operation"`
	Num1      float64 `json:"Num1"`
	Num2      float64 `json:"Num2"`
}

type ResponseData struct {
	Operation string  `json:"operation"`
	Result    float64 `json:"result"`
}

var data Data
var responseData ResponseData

func main() {

	http.HandleFunc("/calculator", handleCalculate) // endpoint
	fmt.Println("Start Server......")
	http.ListenAndServe("localhost:8080", nil) // listen and serve
}

func handleCalculate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			errorHandling(w, err)
			http.Error(w, "Error decoding data", http.StatusBadRequest)
			return
		}
		responseData.Operation = data.Operation

		switch data.Operation {
		case "sum":
			responseData.Result = calculateSum(data)
		case "subtract":
			responseData.Result = calculateSubtract(data)
		case "multiplication":
			responseData.Result = calculateMultiplication(data)
		case "division":
			responseData.Result = calculateDivision(data)
		default:
			http.Error(w, "Invalid operation", http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(responseData)

	}
}

func errorHandling(w http.ResponseWriter, err error) {
	fmt.Println(err)
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}

func calculateSum(data Data) float64 {
	fmt.Println("SUM")
	return data.Num1 + data.Num2
}

func calculateSubtract(data Data) float64 {
	return data.Num1 - data.Num2
}

func calculateMultiplication(data Data) float64 {
	return data.Num1 * data.Num2
}

func calculateDivision(data Data) float64 {
	return data.Num1 / data.Num2
}
