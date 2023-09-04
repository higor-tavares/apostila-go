package main
import "fmt"

func main() {
	c := make(chan string, 3)
	go write(c)
	consumes(c)
}
//bidirectional
//func write(c chan string) {
func write(c chan<- string) {
	c <- "hello"
	c <- "buffered"
	c <- "channels with close!"
	close(c)
}
//bidirectional
//func consumes(c chan string) {
func consumes(c <-chan string) {
	for value := range c {
		fmt.Println(value)
	}
}