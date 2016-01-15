package main

import "fmt"

func main2() {

	// We'll iterate over 2 values in the `queue` channel.
	queue := make(chan string)
	queue <- "one"
	queue <- "two"
	close(queue)

	// This `range` iterates over each element as it's
	// received from `queue`. Because we `close`d the
	// channel above, the iteration terminates after
	// receiving the 2 elements. If we didn't `close` it
	// we'd block on a 3rd receive in the loop.
	for elem := range queue {
		fmt.Println(elem)
	}
}

