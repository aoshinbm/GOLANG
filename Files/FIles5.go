package main

import (
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func writeDataInFile(fName string) {
	fHndlr, err := os.Create(fName)
	check(err)

	barr := []byte{96, 97, 98}
	nb, err := fHndlr.Write(barr)
	check(err)
	fmt.Println("No of bytes: ", nb)

	nb, err = fHndlr.Write(string)
}
func main() {

}
