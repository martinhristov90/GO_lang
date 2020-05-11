package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Starting countdown")
	tick := time.Tick(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick // Blocking for 1 second, discarding the returd timestamp
	}

}
