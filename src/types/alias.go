package main
import (
	"fmt"
)
//create alias to slice of strings
type ProductList []string

func main() {
	products := make(ProductList, 5)
	products[0] = "Book Go in Action"
	products[1] = "Raspberry PI pico"
	products[2] = "Mouse"
	products[3] = "Keyboard"
	products[4] = "Cat food"
	fmt.Println(products)
	books, eletronics, anothers := products.Categorize()
	fmt.Println("Eletronics: ", eletronics)
	fmt.Println("Books: ", books)
	fmt.Println("Anothers : ", anothers)
}
//extends funcionality from a type. For example: []string 
func (list ProductList) Categorize() ([]string, []string, []string) {
	var books, eletronics, anothers []string
	for _, p := range list {
		switch p {
		case "Mouse", "Keyboard", "Raspberry PI pico":
			eletronics = append(eletronics, p)
		case "Book Go in Action":
			books = append(books, p)
		default:
			anothers = append(anothers, p)
		}
	}
	return books, eletronics, anothers
}