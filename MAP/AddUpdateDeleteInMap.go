package main

import "fmt"

func main() {
	CODES := map[string]string{"en": "English", "fr": "French", "hi": "Hindi"}
	fmt.Println(CODES)

	//adding a new item to the map using a new index key and assigning a value to it
	CODES["it"] = "Italian"
	fmt.Println(CODES)

	//update the value of specific item by referring to its key name
	CODES["en"] = "English Language"
	//if adding an already existing key then it will simply override or update its value
	fmt.Println(CODES)

	//to delete a key-value pair from map
	delete(CODES, "en") //(name_of_map,key_name)
	fmt.Println(CODES)

	//looping or iterating over a map
	for key, value := range CODES {
		fmt.Println(key, "==>>", value)
	}

	//Truncate  a map
	//Truncate means clearing all elements from map
	//1st method to truncate a map is:= iterating over map and deleting one by one
	for key, _ := range CODES {
		delete(CODES, key)
	}
	fmt.Println(CODES)
	//2nd method to truncate a map is:= to reinitialize it with empty map
	CODES = make(map[string]string)
	fmt.Println(CODES)
}
