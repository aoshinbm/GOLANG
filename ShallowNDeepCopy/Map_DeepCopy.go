// Write a program to deep copy a map in Go.
package main

import "fmt"

func main() {

	traveller := make(map[string]string)
	traveller["1st"] = "Nagaland"
	traveller["2nd"] = "Turkey"
	traveller["3rd"] = "Sikkim"
	traveller["4th"] = "Dubai"
	traveller["5th"] = "Rajasthan"

	fmt.Println("-----------------------------------")
	fmt.Println("Original Map: ", traveller)

	//toVisit := make(map[string]string)
	/*for count, place := range traveller {
		toVisit[count] = place
	}*/

	toVisit := traveller
	toVisit["4th"] = "Japan"

	fmt.Println("-----------------------------------")
	fmt.Println("Original Map: ", &traveller)
	fmt.Println("-----------------------------------")
	fmt.Println("Copied Map: ", toVisit)
	fmt.Println("-----------------------------------")

}
