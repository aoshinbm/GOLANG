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
	player1 := Player{"Sachin Tendulkar", "1989-2012", 463, 18426, 44.83}
	player2 := Player{"Saurav Ganguly", "1992-2007", 308, 11221, 40.95}
	player4 := Player{"Rahul Dravid", "1996-2011", 340, 10768, 39.15}
	players := []Player{
		player1,
		player2,
		player4,
	}
	for _, plyr := range players {
		fmt.Println(plyr)
	}
	fmt.Println("")
	fmt.Printf("\nConverting slice of structs to JSON \n")
	fmt.Printf("Used Pretty printing \n")
	bytesStructSlice, _ := json.MarshalIndent(players, " ", "\t")
	json.MarshalIndent(players, " ", "\t")
	fmt.Println(string(bytesStructSlice))

}
