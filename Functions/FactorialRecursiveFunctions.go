package main

import "fmt"

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
func main() {
	n := 7
	result := factorial(n)
	fmt.Println("Factorial of ", n, " is : ", result)
}

/*
factorial(7)
!=1
returns 7*factorial(6)=7*720	=5040
  returns 6*factorial(5)=6*120=720
    returns 5*factorial(4)=5*24=120
	    returns 4*factorial(3)=4*6=24
    		returns 3*factorial(2)=3*2=6
    			returns 2*factorial(1)=2*1=2
	    			returns 1=1

*/
