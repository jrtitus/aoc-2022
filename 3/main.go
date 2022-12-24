package main

import (
	"fmt"
	"os"
	"strings"
)

// I hate this function so much
func findCommonItem(sacks ...string) rune {
	for _, aChar := range sacks[0] {
		for _, bChar := range sacks[1] {
			for _, cChar := range sacks[2] {
				if aChar == bChar && bChar == cChar {
					return aChar
				}
			}
		}
	}
	return 0
}

func main() {
	// read file and load rucksacks
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	sacks := strings.Split(string(dat), "\n")

	var sum int32
	grp := make([]string, 0)

	for _, sack := range sacks {
		grp = append(grp, sack)

		if len(grp) == 3 {
			common := findCommonItem(grp...)
			sum += convert(common)
			// new group
			grp = make([]string, 0)
		}
	}

	fmt.Printf("the sum of all common items is %v", sum)
}

func convert(charCode rune) int32 {
	// lowercase: priority starts at 1
	if charCode > 96 && charCode < 123 {
		return charCode - 96
	}
	// uppercase: priority starts at 27
	if charCode > 64 && charCode < 91 {
		return charCode - 38
	}
	return 0
}
