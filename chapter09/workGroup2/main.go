package main

import (
	"fmt"
	"sync"
	"time"
)

// counting semaphore
var tokens = make(chan struct{}, 20)

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Started goroutine with id : %d\n", id)
	time.Sleep(10 * time.Second) // Simulating 10 second of work
	wg.Done()                    // The work is done
	<-tokens                     // releases token when the worker is done
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		tokens <- struct{}{} // puts 20 tokens in the channel at a time, when reaches 20 the channel is full and waiting to be drained
		wg.Add(1)
		go worker(i, &wg)
		if len(tokens) == cap(tokens){
			fmt.Println("No more available tokens, waiting for tokens to be freed up")
		}
	}
	wg.Wait()
	fmt.Printf("100 go routines have ended\n")

}
