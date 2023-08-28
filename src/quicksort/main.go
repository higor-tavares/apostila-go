package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1:]
	numbers := make([]int, len(input))
	for i, n := range input {
		number, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("%s is not a number!\n", n)
			os.Exit(1)
		}
		numbers[i] = number
	}
	fmt.Println(quicksort(numbers))
}

func quicksort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}
	numbersCopy := make([]int, len(numbers))
	copy(numbersCopy, numbers)
	beerIndex := len(numbersCopy) / 2
	beer := numbersCopy[beerIndex]
	//remove the element at beerIndex
	numbersCopy = append(numbersCopy[:beerIndex], numbersCopy[beerIndex+1:]...)
	lessThan, greatherThan := partition(numbersCopy, beer)
	return append(append(quicksort(lessThan), beer), quicksort(greatherThan)...)
}

func partition(numbers []int, beer int) (lessThan, greatherThan []int) {
	for _, n := range  numbers {
		if n < beer {
			lessThan = append(lessThan, n)
		} else {
			greatherThan = append(greatherThan, n)
		}
	}
	return
}