package main

import (
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte("<h1 style='color:red'>HELLO HTML WORLD</h1>"))
}
func main() {
	//need to create and register request handler
	//endpoint:= website.com/endpoint
	/*	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
			//endpoint hello will be called
			//calling accorging to HTTP request (GET/POST)
			//response writer is the server will respond with from any request

			rw.Write([]byte("HElloo"))
		}) //inline function declaration
	*/
	http.HandleFunc("/hello", Hello)
	log.Fatal(http.ListenAndServe(":8081", nil))

}
