package main
import "fmt"
func main() {
	numbers := []int{1,2,3,4,5,6}
	double := make([]int, len(numbers))
	copy(double, numbers)
	for i := range double {
		double[i] *= 2
	}
	fmt.Println("Original :", numbers)
	fmt.Println("Copy: ", double)
}