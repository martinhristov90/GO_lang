package main

import(
	"fmt"
)

func main(){

	c := make(chan string, 2) 
	// If the channel is unbuffered this would result in deadlock
	// Buffered channels do not block until they are "full", so in this case when there are two strings inside c, on the third item it will block

	c <- "hello"
	msg := <- c
	fmt.Println(msg)
	
	c <- "world"
	msg = <- c
	fmt.Println(msg)

}