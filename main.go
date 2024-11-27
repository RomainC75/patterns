package main

import (
	"context"
	"design-patterns/math"
	"sync"
	"time"
)

func generator(min int64, max int64, wg *sync.WaitGroup, ctx context.Context) (chan int64, chan int) {

	out := make(chan int64)
	done := make(chan int)
	go func() {
		for i := min; i < max; i++ {
			select {
			case <-ctx.Done():
				break
			default:
				out <- i
			}
		}
		done <- 1
		wg.Done()
	}()
	return out, done
}

func splitter(chanNumber int, in chan int64, gof func(chan int64, chan int64)) []chan int64 {
	out := []chan int64{}
	for i := 0; i < chanNumber; i++ {
		chn := make(chan int64)
		out = append(out, chn)
		gof(in, chn)
	}
	return out
}

func analyser(in chan int64, out chan int64) {
	go func() {
		for {
			select {
			case num := <-in:
				if math.IsPrime(num) {
					out <- num
				}
			}
		}
	}()
}

func receiver(in []chan int64, done chan int) {
	for _, ch := range in {
		go func() {
			for {
				select {
				case <-done:
					return
				case num := <-ch:
					num++
					// fmt.Println("--->", num)
				}
			}
		}()
	}
}

func parallel(start int64, end int64) {
	var wg sync.WaitGroup
	wg.Add(1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	gen, done := generator(1_000_000, 1_500_000, &wg, ctx)

	chans := splitter(15, gen, analyser)

	receiver(chans, done)

	wg.Wait()

}

func serie(start int64, end int64) {
	var i int64
	for i = start; i < end; i++ {
		if math.IsPrime(i) {
			// fmt.Println("--->", i)
		}
	}
}

func main() {

}
