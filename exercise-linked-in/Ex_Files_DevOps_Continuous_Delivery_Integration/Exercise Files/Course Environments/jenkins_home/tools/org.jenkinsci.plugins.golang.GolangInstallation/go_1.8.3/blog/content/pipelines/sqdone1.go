// +build OMIT

package main

import (
	"fmt"
	"sync"
)

// gen sends the values in nums on the returned channel, then closes it.
func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out
}

// sq receives values from in, squares them, and sends them on the returned
// channel, until in is closed.  Then sq closes the returned channel.
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// merge receives values from each input channel and sends them on the returned
// channel.  merge closes the returned channel after all the input values have
// been sent.
func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed or it receives a value
	// from done, then output calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			select {
			case out <- n:
			case <-done: // HL
			}
		}
		wg.Done()
	}
	// ... the rest is unchanged ...

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// Consume the first value from output.
	done := make(chan struct{}, 2) // HL
	out := merge(done, c1, c2)
	fmt.Println(<-out) // 4 or 9

	// Tell the remaining senders we're leaving.
	done <- struct{}{} // HL
	done <- struct{}{} // HL
}
