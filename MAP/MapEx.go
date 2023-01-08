package main

import "fmt"

func main() {

	//creates a new map and adds three pairs to it.
	benelux := make(map[string]string)
	benelux["be"] = "Belgium"
	benelux["nl"] = "Netherlands"
	benelux["lu"] = "Luxembourgh"

	fmt.Println(benelux)
	fmt.Printf("%q\n", benelux)
}
