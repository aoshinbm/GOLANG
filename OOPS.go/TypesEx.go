package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// (p Point) is lyk receiver
// PrintPoint is function_name
// Why in PrintPoint "P" is capital is bcuz of accesssing from the package
func (p Point) PrintPoint() {
	fmt.Println(p.x, p.y)
}

// (p Point) is lyk receiver
// DistToOrigin ==> D is capital letter is bcuz of access
func (p Point) DistToOrigin() float64 {
	//	temp := (p.x * p.x) + (p.y * p.y)
	//	fmt.Println(temp)
	powerr := math.Pow(p.x, 2) + math.Pow(p.y, 2)
	fmt.Println(powerr)
	dist := math.Sqrt(powerr)

	return dist
	//	return math.Sqrt(math.Pow(p.x, 2) + math.Pow(p.y, 2))
}

// func (p Point) Offset(offsetValue float64) {
// here v r deferencing
// *Point ==> is called pointer receiver
func (p *Point) Offset(offsetValue float64) {
	p.x = p.x + offsetValue
	p.y = p.y + offsetValue
	fmt.Println(p.x, p.y) //here value is incremented
}
func (p *Point) InitialValue(x, y float64) {
	p.x = x
	p.y = y
}

// wen big data is shud be passed to function the pass by address is gud
// wen small data is passed to function the pass by value is gud , change will be in local storage
func main() {
	p1 := Point{3, 4}
	p1.PrintPoint()

	fmt.Println(p1.DistToOrigin())

	p1.Offset(2)
	p1.PrintPoint() //but it prints the original vaues only

	var p2 Point
	p2.InitialValue(2.0, 3.2)
	p2.PrintPoint()
}
