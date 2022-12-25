package main

import "fmt"

func main() {
	var stateCapitals = make(map[string]string)

	stateCapitals["MAHA"] = "MUMBAI"
	stateCapitals["TamilNADU"] = "Chennai"
	stateCapitals["KERALA"] = "TRIVANDRUM"
	stateCapitals["XYZ"] = "ABC"
	stateCapitals["KARNATAK"] = "BENGALURU"
	for state, capital := range stateCapitals {
		fmt.Printf("Capital of %s is %s \n", state, capital)
	}

	fmt.Println()
	stateCapitals["J&K"] = "SRINAGAR"
	stateCapitals["RAJASTHAN"] = "JAIPUR"
	for state, capital := range stateCapitals {
		fmt.Printf("Capital of %s is %s \n", state, capital)
	}

	fmt.Println()
	delete(stateCapitals, "XYZ")
	for state, capital := range stateCapitals {
		fmt.Printf("Capital of %s is %s \n", state, capital)
	}

	x := map[string]int{
		"ian": 1, "harris": 2}
	for i, j := range x {
		if i == "harris" {
			fmt.Print(i, j)
		}
	}

}
