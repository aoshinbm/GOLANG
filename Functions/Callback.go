package main

import "fmt"

// Define a function that takes a callback as an argument
func doWork(callback func()) {
    // Perform some work
  fmt.Println("Doing work...")
  calling this  https://reqres.in/api/users?page=2
  data in datavariable
  fmt.Println("I am calling IMDB database rest api.v  )

    // Call the callback
    callback(data)
}

func main() {
    // Define a function that will be passed as a callback
    done := func() {
        fmt.Println("Work is done!")
    }

    // Call the doWork function and pass the done function as a callback
    doWork(done) // Output: "Doing work...", "Work is done!"
}
