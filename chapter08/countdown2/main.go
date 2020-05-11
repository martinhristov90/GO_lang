package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	fmt.Println("Starting countdown")
	tick := time.Tick(1 * time.Second)
	abort := make(chan struct{}, 0)

	go func() {
		os.Stdin.Read(make([]byte, 1)) // Reads one byte from the stdin

		abort <- struct{}{}

	}()

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick: // Blocking for 1 second, discarding the returd timestamp
		case <-abort:
			fmt.Println("Aborting")
			return
		}

	}

	// Adding abort function

}
