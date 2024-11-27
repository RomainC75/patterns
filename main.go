package main

import (
	"fmt"
	"sync"
)

func generator(min int64, max int64, wg *sync.WaitGroup) (chan int64, chan int) {
	out := make(chan int64)
	done := make(chan int)
	go func() {
		for i := min; i < max; i++ {
			out <- i
		}
		done <- 1
		wg.Done()
	}()
	return out, done
}

func receiver(in chan int64, done chan int) {
	go func() {
		for {
			select {
			case <-done:
				return
			case num := <-in:
				fmt.Println("--->", num)
			}
		}
	}()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	gen, don := generator(1, 1_000_000, &wg)
	receiver(gen, don)

	wg.Wait()
}
