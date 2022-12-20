package main

import (
	"fmt"
	"os"
	"strings"
)

type Compartment struct {
	items string
}

func newCompartment(items string) Compartment {
	return Compartment{items: items}
}

func (c Compartment) findDuplicate(alt Compartment) rune {
	for _, char := range c.items {
		for _, altChar := range alt.items {
			if char == altChar {
				return char
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
	// items in each compartment
	for _, sack := range sacks {
		c1, c2 := splitItems(sack)
		dup := c1.findDuplicate(c2)

		sum += convert(dup)
	}

	fmt.Printf("the sum of all duplicate items is %v", sum)
}

func splitItems(input string) (Compartment, Compartment) {
	half := len(input) / 2
	return newCompartment(input[:half]), newCompartment(input[half:])
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
