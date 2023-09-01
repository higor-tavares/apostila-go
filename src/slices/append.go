package main
import "fmt"

func main() {
	slice := []int {2,3,4,5,6,7,9}
	fmt.Println("Slice :", slice)
	slice = append([]int{1}, slice...)
	fmt.Println("Slice with value append on first position :", slice)
}