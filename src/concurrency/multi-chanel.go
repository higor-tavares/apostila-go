package main
import "fmt"

func selector(numbers []int, even, odd chan<- int, ready chan<- bool) {
	for _, n := range numbers {
		if n%2	 == 0 {
			odd <- n
		} else {
			even <- n
		}
	}
	ready <- true
}

func main() {
	e, o := make(chan int), make(chan int)
	r := make(chan bool)
	n := []int{1,23,42,5,8,6,7,4,99,100}
	go selector(n, e, o,r)
	var even, odd []int
	ready := false
	for !ready {
		select {
		case n := <- e:
			even = append(even, n)
		case n := <- o:
			odd = append(odd, n)
		case ready = <-r:
		}
	}
	fmt.Printf("Even: %v | Odd: %v\n", even, odd)
}