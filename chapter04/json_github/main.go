package main

import (
	"fmt"
	"github"
	"log"
	"os"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatalf("%v\n",err)
	}

	fmt.Printf("%d Issues\n", result.TotalCount)

	for _, item := range result.Items {
		fmt.Printf("#%-5d %d %9.9s %.55s\n", item.Number, item.Id, item.User.Login, item.Title,)
	}
}
