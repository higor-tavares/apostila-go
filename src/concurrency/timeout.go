package main
import (
	"fmt"
	"time"
)
func exec(c chan<- bool) {
	time.Sleep(5*time.Second)
	c <- true
} 
func main() {
	c := make(chan bool, 1)
	go exec(c)
	fmt.Println("waiting...")
	end := false
	for !end {
		select {
		case end = <-c:
			fmt.Println("task sucesssfuly completed!")
		case <-time.After(2*time.Second):
			fmt.Println("error: timeout during task execution!")
			end = true
		}
	}
}