package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan string)
	go count("sheep", c)

	for {
		msg, open := <-c // Waiting to recieve a msg on c forever, results in deadlock if c is not closed
		if !open { // If the channel is closed, break out of the for loop and finish main
			break
		}
		fmt.Println(msg)
	}
	
	// nicer way to do the above
	//for msg := range c {
	//	fmt.Println(msg)
	//}

}

func count(str string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- fmt.Sprintf("i at the moment: %d", i)
		time.Sleep(1 * time.Second)
	}

	close(c) // Closing the channel when there is nothing left to be sent.If do not close the channel, causes a deadlock on the last iteration. This is due that main goroutine expects to receive values forever, but nothing is sent on the c channel.
}
