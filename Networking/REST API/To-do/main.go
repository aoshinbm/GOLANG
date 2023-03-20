package main

import (
	"github.com/gin-gonic/gin"
)

// how To-Do structure will look is defined by "struct"
type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

// array
var todo = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Wash Utensils", Completed: false},
	{ID: "3", Item: "Mow the lawn", Completed: false},
	{ID: "4", Item: "Empty Dustbin", Completed: false},
}

func main() {

	router := gin.Default() //creates a server
	router.Run()

}
