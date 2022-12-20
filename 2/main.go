package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// read file and load games
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

func getWinningMove(opp string) int {
	if opp == "A" {
		return 2
	}
	if opp == "B" {
		return 3
	}
	return 1
}

func getLosingMove(opp string) int {
	if opp == "A" {
		return 3
	}
	if opp == "B" {
		return 1
	}
	return 2
}

func getDrawMove(opp string) int {
	if opp == "A" {
		return 1
	}
	if opp == "B" {
		return 2
	}
	return 3
}

func calculateGame(g []string) int {
	// A B C
	opp := g[0]
	// X - Lose; Y - Draw; Z - Win
	outcome := g[1]

	// Win condition
	if outcome == "Z" {
		return 6 + getWinningMove(opp)
	}

	// Loss condition
	if outcome == "X" {
		return getLosingMove(opp)
	}

	// Draw condition
	return 3 + getDrawMove(opp)
}
