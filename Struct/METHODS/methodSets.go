package main

import "fmt"

type Student struct {
	name  string
	grade []int //slice
}

// method has a receiver of Student type pointer
func (s *Student) displayName() {
	fmt.Println(s.name)
}

func (s *Student) calculatePercentage() float64 {
	sum := 0
	for _, v := range s.grade {
		sum += v
	}
	return float64(sum*100) / float64(len(s.grade)*100)
}
func main() {
	s := Student{name: "OLIVIA", grade: []int{35, 49, 90, 65}}
	s.displayName()
	fmt.Printf("%.2f", s.calculatePercentage())
}
