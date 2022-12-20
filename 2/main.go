package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// read file and load caloriesPerElf
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	games := strings.Split(string(dat), "\n")

	var totalScore int
	for _, game := range games {
		g := strings.Split(game, " ")
		totalScore += calculateGame(g)
	}

	fmt.Printf("total score is %d", totalScore)
}

var scores map[string]int = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func calculateGame(g []string) int {
	// A B C
	opp := g[0]
	// X Y Z
	me := g[1]

	// Win condition
	if opp == "A" && me == "Y" || opp == "B" && me == "Z" || opp == "C" && me == "X" {
		return 6 + scores[me]
	}

	// Loss condition
	if opp == "A" && me == "Z" || opp == "B" && me == "X" || opp == "C" && me == "Y" {
		return scores[me]
	}

	// Draw condition
	return 3 + scores[me]
}
