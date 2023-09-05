package main
import (
	"fmt"
	"math"
	"sync"
	"time"
	"runtime"
)

func calc(base float64, control * sync.WaitGroup) {
	defer control.Done()
	n := 0.0
	for i := 0; i < 100000000; i++ {
		n += base / math.Pi * math.Sin(2)
	}
	fmt.Println(n)
}
func main() {
	//use all cores
	//runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GOMAXPROCS(3)
	begin := time.Now()
	var control sync.WaitGroup
	control.Add(3)
	go calc(9.37, &control)
	go calc(6.94, &control)
	go calc(42.47, &control)
	control.Wait()
	fmt.Println("The calc spends %s.", time.Since(begin))
}