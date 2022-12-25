package main

import "fmt"

func main() {
	city := "LONDON"
	{
		country := "UK"
		fmt.Println(city)
		fmt.Println(country)
	}
	fmt.Println(city)
	//	fmt.Println(country)
	/*OUTPUT:
	.\VariableScope.go:13:14: undefined: country

	Variable Scope:- Outer block(func main()) cannot access inner block({}-inside curly brackets)
	variables and vice versa is possible
	*/
}
