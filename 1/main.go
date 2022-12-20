package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cast"
)

var caloriesPerElf []int = make([]int, 0)
var max int = 0

func appendSetMax(sum int) {
	// store sum
	caloriesPerElf = append(caloriesPerElf, sum)
	// determine max
	if sum > max {
		max = sum
	}
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

	fmt.Printf("there are a total of %d elves\n", len(caloriesPerElf))
	fmt.Printf("max calories carried is %d", max)
}
