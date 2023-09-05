package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	begin := time.Now()
	rand.Seed(begin.UnixNano())
	var control sync.WaitGroup
	for i := 0; i < 5; i++ {
		control.Add(1)
		go exec(&control)
	}
	control.Wait()
	fmt.Printf("Execution time: %s.\n", time.Since(begin))
}

func exec(c *sync.WaitGroup) {
	defer c.Done()
	duration := time.Duration(1+rand.Intn(5)) * time.Second
	fmt.Printf("Sleeping for %s...\n",duration)
	time.Sleep(duration)
}