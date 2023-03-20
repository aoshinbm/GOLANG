package main

import (
	"fmt"
	"net/http"
)

func main() {
	// create client
	req, _ := http.NewRequest("GET", "http://localhost:8082", nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	fmt.Println(res.StatusCode)
}
