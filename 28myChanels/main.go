package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("CHAnNELS")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	//receive ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myCh

		fmt.Println(isChannelOpen)
		fmt.Println(val)
		wg.Done()
	}(myCh, wg)

	//send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		// myCh <- 6
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
