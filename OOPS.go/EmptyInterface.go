package main

import "fmt"

func main() {
	person := make(map[string]interface{}, 0)

	person["name"] = "Aoshin"
	person["age"] = 25
	person["height"] = 5.6

	fmt.Printf("%+v", person)
}
