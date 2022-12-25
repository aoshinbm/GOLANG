package main

import "fmt"

func main() {

	arrayLiteral := [6]int{20, 25, 35, 40, 60, 99}
	fmt.Println(arrayLiteral)

	sliceCreated := arrayLiteral[2:5]
	fmt.Println("Slice values are : ", sliceCreated)
	fmt.Printf("Length of slice = %d \nCapacity of slice = %d", len(sliceCreated), cap(sliceCreated))

	arrayString := [4]string{"apple", "peaches", "raspberry", "pear"}
	fmt.Println(arrayString)

	topFootballPlayers := [...]string{"Messi", "LukaModric", "Ronaldo", "Neymar", "Cavani", "CarlesPuyol"}
	fmt.Println(topFootballPlayers)

	slice1 := arrayString[2:3]
	slice2 := arrayString[3:4]
	fmt.Println(slice1, slice2)

	strngArr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	sl1 := strngArr[1:3]
	sl2 := strngArr[2:5]
	fmt.Println(sl1, sl2)

	nilSlice := make([]int, 0)
	fmt.Println(nilSlice)

	x := [...]int{40, 80, 50}
	y := x[0:2]
	z := x[1:3]
	y[0] = 10
	z[1] = 30
	fmt.Print(x)

	/*	x := [...]int{15, 19, 3, 40, 50}
		y := x[0:2]
		z := x[1:4]
		fmt.Print(len(y), cap(y), len(z), cap(z))
	*/
	x := []int{3, 7, 9, 5}
	y := -1
	for _, element := range x {
		if element > y {
			y = element
		}
	}
	fmt.Print(y)
}
