package main

import "fmt"

func main() {
		const name = "Harry Styles"
		const is_muggle = false
		const age = 20

		fmt.Printf("%v : %T \n ", name, name)
		fmt.Printf("%v : %T \n ", is_muggle, is_muggle)
		fmt.Printf("%v : %T \n ", age, age)
	

	fmt.Println("--------------------------------------------------------------------")
	const fname = "Harry Styles"
	fmt.Printf("%v : %T \n ", fname, fname)
	fname = "Monica" //cannot assign to fname (untyped string constant "Harry Styles")
	fmt.Printf("%v : %T \n ", fname, fname)
	/*The compiler is signaling us that you cannot assign
	or reassign a variable because this is a constant variable.*/

	fmt.Println("--------------------------------------------------------------------")
	/*Note:- You cannot declare constant and not initialize it a value 
	and maybe try to assign it a value later on.*/
	const str string
	str="Hermonie Walker"
	fmt.Println(str)
	//ERROR:- "Missing value in constant declaration."

	//You have to assign and declare constant at the same time.

	fmt.Println("--------------------------------------------------------------------")
	const vote := "yes"
	fmt.Println(vote)
	/*The short and variable declaration does not work for constants.
	The compiler will give Error if we try to use the short hand variable assignment operator.
	ERROR: "Unexpected :=" syntax error: unexpected :=, expecting = "*/
}
