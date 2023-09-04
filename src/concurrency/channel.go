package main
import "fmt"

func main() {
	c := make(chan int)
	go product(c)
	value := <- c
	fmt.Println(value)
}

func product(c chan int) {
	c <- 27
}