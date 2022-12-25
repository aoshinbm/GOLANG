// Golang program to read and write the files
package main

// importing the packages
import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Player struct {
	name, period             string
	matchesPlayed, runScored int
	avgScore                 float64
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func readCSVFile(fName string) {
	fhandler, err := os.Open(fName)
	check(err)
	defer fhandler.Close()

	reader := csv.NewReader(fhandler)

	headerr, err := reader.Read()
	fmt.Println(headerr)

	records, err := reader.ReadAll()
	check(err)
	fmt.Println("Printing Records\n")

	var players []Player

	for _, record := range records {
		matchesPlayed, _ := strconv.Atoi(record[2])
		runScored, _ := strconv.Atoi(record[3])
		avgScore, _ := strconv.ParseFloat(record[4], 64)

		record, err := reader.ReadAll()
		check(err)
		fmt.Printf("\nPrinting records having type of %T \n", record)
		fmt.Println(record)

		partData := Player{
			name:          record[0],
			period:        record[1],
			matchesPlayed: matchesPlayed,
			runScored:     runScored,
			avgScore:      avgScore,
		}
		players = append(players, partData)
	}
}

func main() {

	readCSVFile("")
}
