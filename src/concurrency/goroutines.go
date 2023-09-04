package main
import (
	"fmt"
	"time"
)

func sleep() {
	fmt.Println("Goroutine sleeping for 5 seconds...")
	time.Sleep(5*time.Second)
	fmt.Println("Goroutine finished with success!")
}

func main() {
	//call go routine tha depends from main 
	//this code this will never end
	go sleep()
	fmt.Println("Main sleep for 3 seconds...")
	time.Sleep(3*time.Second)
	fmt.Println("Main finished with success!")
}