package main

import "fmt"

func main() {
	codes := make(map[string]string)
	fmt.Println(codes) //an empty map will be printed
	/*
		OUTPUT:-
		map[]
	*/

	//to determine items or key-value pair in a map
	CODES := map[string]string{"en": "English", "fr": "French", "hi": "Hindi"}
	fmt.Println(len(CODES))

	//access items in map by referring key value
	fmt.Println(CODES["en"])
	fmt.Println(CODES["fr"])
	fmt.Println(CODES["hi"])

	//get a key in the map
	code := map[string]int{"en": 1, "fr": 2, "hi": 3}
	fmt.Println(code)
	value, found := code["en"]
	fmt.Println(found, value) //true 1 (bcuz this key exists in map &
	//true is boolean found and 1 is the value associated with it)
	value, found = code["hh"]
	fmt.Println(found, value) //false 0 (bcuz this key doesnt exists in map)

}
