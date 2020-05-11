package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Hello from inside c1 channel, every 500 ms"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c2 <- "Hello from inside c2 channelm every 2s"
			time.Sleep(2 * time.Second)
		}
	}()

	//for {
	//	fmt.Println(<-c1)
	//	fmt.Println(<-c2)
	//	// In this case the output is shown every 2s, because fmt.Println(<-c2) blocks for two seconds.
	//}

	for {
		select { // picks s channel that is available
			case msg1 := <-c1:
				fmt.Println(msg1) // msg every 500 ms 
			case msg2 := <-c2:
				fmt.Println(msg2) // msg every 2s
		}
		//Hello from inside c1 channel, every 500 ms
		//Hello from inside c1 channel, every 500 ms
		//Hello from inside c1 channel, every 500 ms
		//Hello from inside c1 channel, every 500 ms
		//Hello from inside c2 channelm every 2s
	}

}
