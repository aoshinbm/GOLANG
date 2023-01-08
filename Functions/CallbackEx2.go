package main

import "fmt"

// Define a function that takes a callback as an argument
func normalAPICall() {
    // Perform some work
  fmt.Println("Doing work...")
  calling this api  https://reqres.in/api/users?page=2   returns json data           /// it is taking 1 mins
  storing the data in datavariable
  
// On execution datavariable = null.

  // Main logic starts here for our program
	
 fmt.Println(datavariable[1].firstname)   // Index out of bound
 fmt.Println(datavariable[1].lastname)    // Index out of bound

}
