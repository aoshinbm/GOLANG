package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, myArray)
	myByteSlice := buf.Bytes()
	fmt.Println(myByteSlice)
}
