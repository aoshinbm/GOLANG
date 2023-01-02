package main

import (
	"fmt"
	"os"
	"reflect"
)

/*
Command-line arguments are options and data that are passed to programs.
We usually pass arguments to console programs,
but sometimes we pass arguments to GUI programs as well.

The os.Args holds the command-line arguments.
The first value in this slice is the name of the program,
while the os.Args[1:] holds the arguments to the program.
The individual arguments are accessed with indexing operation.
*/
func main() {

	//We get and print the first argument, which is the program name
	prg_name := os.Args[0]
	fmt.Printf("The program name is %s\n", prg_name)

	//We get all the received arguments
	names := os.Args[1:]

	//print the type which holds the arguments (slice)
	fmt.Println(reflect.TypeOf(names))

	for _, name := range names {
		//go through the arguments and build a message from each of them
		fmt.Printf("Hello, %s!\n", name)
	}
}
