package main

import (
	"context"
	"fmt"
	"sync"
)

func fanIn(ctx context.Context, fetchers ...<-chan interface{}) <-chan interface{} {

	fmt.Println(" fanIn starts")

	combinedFetcher := make(chan interface{})

	// 1

	var wg sync.WaitGroup
	wg.Add(len(fetchers))

	fmt.Println("3 go routines waits ")

	// 2

	for _, f := range fetchers {

		fmt.Println(" go routines ranges", f)

		f := f

		go func() {

			// 3

			fmt.Println("unnamed go routines executes")

			defer wg.Done()
			for {
				select {
				case res := <-f:
					fmt.Println(" data copied from one channel to another")
					combinedFetcher <- res

				case m := <-ctx.Done():
					fmt.Println(" context exited goroutine", m)
					return
				}
			}
		}() //registering the 3 gorouitne
	}

	// 4
	// Channel cleanup

	fmt.Println("2nd unnamed go routines may register")

	go func() {

		fmt.Println(" waiting for all goroutines")
		wg.Wait()
		close(combinedFetcher)

		fmt.Println(" waiting is over")
	}()

	return combinedFetcher
}

func main() {

	fmt.Println(" main starts")

	ctx := context.TODO()

	ch1 := make(chan interface{})

	ch2 := make(chan interface{})

	ch3 := make(chan interface{})

	ch4 := fanIn(ctx, ch1, ch2, ch3) // 4 go routines

	fmt.Println(" ch1 written")

	ch1 <- "coderrange"

	fmt.Println(" ch2 written")

	ch2 <- "glolang"

	fmt.Println(" ch3 written")

	ch3 <- "rust"

	fmt.Println(" main ends", <-ch4)

}