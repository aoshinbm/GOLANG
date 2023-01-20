package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	u1 := User{1, "Eric", "er@gmail.co"}
	u2 := User{2, "Rose", "rose@gmail.co"}
	u3 := User{3, "Rio", "rio@gmail.co"}
	users := []User{
		u1,
		u2,
		u3,
	}

	//user := &User{ID: 1, Name: "John Smith", Email: "john.smith@example.com"}

	fmt.Println(users)
	fmt.Printf("%T \n", users)
	// serialize the struct to json
	data, _ := json.MarshalIndent(users, " ", "\t")
	fmt.Println(string(data))
	fmt.Printf("\n %T", data)

}
