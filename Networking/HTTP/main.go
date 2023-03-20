package main

import "net/http"

type Welcome string

func main() {
	//server
	server := http.Server{
		//http server is a struct that accepts 2 parameters
		Addr: ":8080",
		//1st is address which is a port that u want to use for server
		Handler: nil,
		//2nd is handler
	}
	//run the server
	server.ListenAndServe()

	/*
		Here the sever will run but it wont do anything bcuz no logic
		or client to server content
		therefore it will jz run on port that mentioned nothing else
	*/
}
