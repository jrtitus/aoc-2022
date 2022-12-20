package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cast"
)

var caloriesPerElf []int = make([]int, 0)

func appendSetMax(sum int) {
	// store sum
	caloriesPerElf = append(caloriesPerElf, sum)
}

func main() {
	// read file and load caloriesPerElf
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	strs := strings.Split(string(dat), "\n")

	var sum int
	for _, str := range strs {
		sum += cast.ToInt(str)
		if str == "" {
			appendSetMax(sum)
			sum = 0
		}
	}
	appendSetMax(sum)

	// Only sort once
	sort.Ints(caloriesPerElf)

	elves := len(caloriesPerElf)
	fmt.Printf("there are a total of %d elves\n", elves)

	topThree := caloriesPerElf[elves-1] + caloriesPerElf[elves-2] + caloriesPerElf[elves-3]
	fmt.Printf("the 3 elves carrying the most calories carry a total of %d", topThree)
}
