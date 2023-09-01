package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	words := os.Args[1:]
	statistics := getStatistics(words)
	print(statistics)
}

func getStatistics(words []string) map[string]int {
	statistics := make(map[string]int)
	for _, w := range words {
		firstLetter := strings.ToUpper(string(w[0]))
		counter, find := statistics[firstLetter]
		if find {
			statistics[firstLetter] =  counter + 1
		} else {
			statistics[firstLetter] = 1
		}
	}
	return statistics
}

func print(statistics map[string]int) {
	for key, value := range statistics {
		fmt.Printf("%s = %d\n", key, value)
	}
}