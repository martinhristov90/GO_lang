package main

import(
	"fmt"
	"flag"
	"time"
)

func main(){

	var period = flag.Duration("period",1 * time.Second, "sleep period")

	flag.Parse()
	fmt.Printf("Sleeping for %v...",*period)
	time.Sleep(*period)
	fmt.Println()

}