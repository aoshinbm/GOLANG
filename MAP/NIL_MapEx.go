package main

import "fmt"

func main() {
	var codes map[string]string //nil map

	//add a key-value pair
	codes["en"] = "English"
	fmt.Println(codes)

	//v annot add key-value pair in nil map

}

/*
OUTPUT:=
panic: assignment to entry in nil map
*/
