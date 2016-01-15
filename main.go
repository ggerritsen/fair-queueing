package main
import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Printf("Start\n")

	c1 := make(chan int, 10000)
	c2 := make(chan int, 10000)
	d1 := make(chan int)

	// source channels
	go func(c chan int) {
		for i := 1; i < 10000; i++ {
			c <- i
		}
		close(c)
	}(c1)
	go func(c chan int) {
		for i := -1; i > -20; i-- {
			c <- i
		}
		close(c)
	}(c2)

	fanIn(d1, c1, c2)

	// sink
	for i := range d1 {
		fmt.Printf("Got %d\n", i)
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Printf("Done\n")
}

func fanIn(out chan int, ins... chan int) {
	for _, in := range ins {
		wg.Add(1)
		go func(out chan int, in chan int) {
			for i := range in {
				out <- i
			}
			wg.Done()
		}(out, in)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
}