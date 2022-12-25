package main

import (
	"fmt"
	"os"
)

func main() {
	testString := "this is the file created for testing"
	//	err := ioutil.WriteFile("test.txt", []byte(testString), 0666)
	err := os.WriteFile("test.txt", []byte(testString), 0666)
	if err != nil {
		panic(err)
	}

	//	data, err := ioutil.ReadFile("test.txt")
	data, err := os.ReadFile("test.txt")
	fmt.Printf("%T\n", data)
	fmt.Println(data)
	fmt.Println(string(data))

}
