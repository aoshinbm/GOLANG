package main

import "fmt"

func main() {
	type Player struct {
		name, position, country string
		age                     int
	}
	player1 := Player{
		name:     "Luka Modric",
		country:  "CROATIA",
		position: "Midfielder",
		age:      37,
	}
	player2 := Player{
		name:     "CR7",
		country:  "PORTUGAL",
		position: "Forward & Captain",
		age:      37,
	}
	player3 := Player{
		name:     "Lionel MESSI",
		country:  "ARGENTINA",
		position: "Forward",
		age:      35,
	}

	fmt.Println("PLAYER 1 : ", player1)
	fmt.Println("PLAYER 2 : ", player2)
	fmt.Println("PLAYER 3 : ", player3)

	fmt.Println()
	playerList := []Player{
		player1,
		player2,
		player3,
	}
	for _, plyr := range playerList {
		fmt.Printf("%s ,%s \n", plyr.name, plyr.country)
	}

	fmt.Println()
	player4 := Player{
		name:     "PELE",
		country:  "BRAZIL",
		position: "Forward",
		age:      82,
	}
}
