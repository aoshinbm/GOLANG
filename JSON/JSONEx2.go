package main

import (
	"encoding/json"
	"fmt"
)

type Player struct {
	Name                      string
	Period                    string
	MatchesPlayed, RunsScored int
	AvgScore                  float64
}

func main() {
	fmt.Println("Creating an variable of Type Player and print the same")
	player1 := Player{"Sachin Tendulkar", "1989-2012", 463, 8426, 44.83}
	fmt.Println(player1)
	fmt.Println("")
	fmt.Println("Converting this variable to json")
	bytes, _ := json.Marshal(player1)
	fmt.Println(string(bytes))
	fmt.Println("------------")

	fmt.Println("Creating a map and prting the same")
	matchesPlayed := make(map[string]int)

	matchesPlayed["Sachin"] = 463
	matchesPlayed["Ganguly"] = 308
	matchesPlayed["Dhoni"] = 347
	matchesPlayed["Kohli"] = 262
	matchesPlayed["Dravid"] = 340
	fmt.Println(matchesPlayed)
	fmt.Println("")
	fmt.Println("***************************")

	fmt.Println("Converting this variable to json")
	bytes, _ = json.Marshal(matchesPlayed)
	fmt.Printf("%T", bytes)
	fmt.Println("***************************")
	fmt.Printf("%T %s", string(bytes), string(bytes))
}
