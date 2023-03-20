package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	user := &User{ID: 1, Name: "John Smith", Email: "john.smith@example.com"}
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(user) // Encode the struct
	var newUser User
	dec := gob.NewDecoder(&buf)
	dec.Decode(&newUser) // Decode the struct
	fmt.Println(newUser)
}
