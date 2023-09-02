package main

import "fmt"

type ProductList []string

func printSlice(slice []string) {
	fmt.Println("Slice: ", slice)
}

func printList(list ProductList) {
	fmt.Println("Product List: ", list)
}

func main() {
	list := ProductList{"Book of Go", "Book of Java", "GitHub CodeSpace Subscription"}
	slice := []string{"Another Book", "Arduino UNO R3"}
	//doing conversions
	printSlice([]string(list))
	printList(ProductList(slice))
}