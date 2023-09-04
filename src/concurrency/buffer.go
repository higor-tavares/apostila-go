package main
import "fmt"

func main() {
	c := make(chan string, 3)
	go write(c)
	//cause deadlock
	//fmt.Println(<-c, <-c, <-c,<-c)
	fmt.Println(<-c, <-c, <-c)
}
func write(c chan string) {
	c <- "hello"
	c <- " buffered "
	c <- "channels with go!"
}