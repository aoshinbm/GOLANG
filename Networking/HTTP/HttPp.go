package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type STOCK struct {
	symbol         string
	companyName    string
	latestPrice    float64
	change         float64
	changePercent  int
	marketCap      int
	avgTotalVolume int
	week52High     float64
	week52Low      float64
}

func main() {
	http.HandleFunc("/Get", handleGetRequest)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func handleGetRequest(writer http.ResponseWriter, request *http.Request) {
	var stock []STOCK
	resp, err := http.Get("https://my-json-server.typicode.com/aoshinbm/samplestockdata/stocks")
	//resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)

	if request.Method == http.MethodPost {
		err := json.NewDecoder(request.Body).Decode(&user)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(writer, "Invalid request: %v", err)
			return
		}
		users = append(users, user)
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
