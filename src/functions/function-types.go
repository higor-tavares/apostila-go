package main
import "fmt"

type Aggregation func(current, aggregated int) int

func Aggregate(
	values []int,
	initialValue int,
	fn Aggregation) int {
	aggregated := initialValue
	for _, current := range values {
		aggregated = fn(current, aggregated)
	}
	return aggregated
}

func Sum(values []int) int {
	sum := func(n, m int) int {
		return n + m
	}
	return Aggregate(values, 0, sum)
}

func Multiply(values []int) int {
	multiply := func(n, m int) int {
		return n * m
	}
	return Aggregate(values, 1, multiply)
}

func main() {
	values := []int{3,-2,5,7,8,22,32,-1}
	fmt.Println(Sum(values))
	fmt.Println(Multiply(values))
}