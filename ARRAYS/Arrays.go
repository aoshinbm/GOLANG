package main

import "fmt"

func main() {

	var testArray [5]int
	fmt.Println(testArray)
	/* OUTPUT:-
	[0 0 0 0 0]
	*/

	testArray[0] = 2
	testArray[1] = 12
	testArray[2] = 42
	fmt.Println(testArray)

	//declaring array with predefined values
	arrayLiteral := [6]int{20, 25, 35, 40, 60, 99}
	fmt.Println(arrayLiteral)

	arrayString := [4]string{"apple", "peaches", "raspberry", "pear"}
	fmt.Println(arrayString)

	arrayFloat := [6]float32{2.10, 25.5, 3.15, 46.0, 0.60, 9.9}
	fmt.Println(arrayFloat)

	topFootballPlayers := [...]string{"Messi", "LukaModric", "Ronaldo", "Neymar"}
	fmt.Println(topFootballPlayers)

	//range function for an array
	nums := [3]int{2, 3, 4}
	for i, num := range nums {
		fmt.Println(i, num)
	}
	num := [3]int{2, 3, 4}
	for _, numbr := range num {
		fmt.Println(numbr)
	}

	frnz := [3]string{"radhu", "dolly", "mona"}
	for i, strr := range frnz {
		fmt.Println(i, strr)
	}
	frnd := [3]string{"radhu", "dolly", "mona"}
	for _, strrn := range frnd {
		fmt.Println(strrn)
	}

	marks := [5]float64{52.3, 35.9, 41.8, 74.6, 85.0}
	fmt.Println(len(marks))
	var sum float64 = 0.0
	var avgerage float64 = 0.0
	for i, m := range marks {
		fmt.Println(i, m)
		sum = marks[i] + marks[i+1]
		avgerage = sum / 5.0
	}
	fmt.Println(sum, avgerage)
}
