package main

import "net/http"

func main() {
	http.HandleFunc("/get", handleGetRequest)
	http.ListenAndServe("127.0.0.1:8000", nil)

}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Welcome to web browser"))
}
