package main
import (
	"fmt"
	"time"
)

func generateFibonnaci(n int) func() {
	return func() {
		a, b := 0,1
		fib := func() int {
			a, b = b, a+b
			return a
		}
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", fib())
		}
	}
}

func analize(function func()) {
	begin := time.Now()
	function()
	fmt.Printf("Execution time: %s \n\n",time.Since(begin))

}

func main() {
	analize(generateFibonnaci(10))
	analize(generateFibonnaci(50))
	analize(generateFibonnaci(20))
}