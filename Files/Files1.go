package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	_, err := os.Stat("words.txt")

	//check if the given file exists or not
	//With os.Stat function on the file.
	//If the function returns the os.ErrNotExist error,
	//the file does not exist.
	if errors.Is(err, os.ErrNotExist) {

		fmt.Println("file does not exist")
	} else {

		fmt.Println("file exists")
	}
}
