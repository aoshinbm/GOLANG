package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//http.HandleFunc is used to handle request
	//it takes 2 parameters: path & function
	//syntax of func is func(w http.ResponseWriter, r *http.Request)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		data, _ := ioutil.ReadAll(r.Body)
		//body is like we can use any method of read to read from that stream
		log.Printf("Data %v ", data)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Adios my dear people")
	})

	//to create a webserver http.ListenAndServe is used
	//it has 2 parameters :address & handler
	//addres is bind address which binds all ip address at port number (Eg :8080)
	//This is how webserver is created
	http.ListenAndServe(":8082", nil)

}
