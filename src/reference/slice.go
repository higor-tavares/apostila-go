package main

import "fmt"

func main() {
	original := []int{1,2,3,4,5}
	fmt.Println("Original: ", original)
	copied := original[:3]
	fmt.Println("Copied: ", copied)
	copied[1] = 13
	fmt.Println("Original after change: ", original)
	fmt.Println("Copied atter change: ", copied)
}