package main

import (
	"fmt"
	"log"
	"os"
)

var cwd string

func init() {
	fmt.Printf("executing init\n")
	var err error
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd has failed: %v", err)
	}

	log.Printf("current dir %v\n", cwd)
}

func main() {

	fmt.Printf("current dir in main func %v\n", cwd)
}
