package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cast"
)

func getLowsAndHighs(pairs []string) (int, int, int, int) {
	firstRange := strings.Split(pairs[0], "-")
	secondRange := strings.Split(pairs[1], "-")

	return cast.ToInt(firstRange[0]), cast.ToInt(firstRange[1]), cast.ToInt(secondRange[0]), cast.ToInt(secondRange[1])
}

func main() {
	// read file and load elfPairs
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	elfPairs := strings.Split(string(dat), "\n")

	var count int
	for _, pair := range elfPairs {
		if pair == "" {
			continue
		}
		pairs := strings.Split(pair, ",")
		l1, h1, l2, h2 := getLowsAndHighs(pairs)
		if l1 >= l2 && h1 <= h2 || l2 >= l1 && h2 <= h1 {
			count += 1
		}
	}

	fmt.Printf("%d pairs contain nested ranges", count)
}
