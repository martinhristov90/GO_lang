package main

import (
	"bufio"
	"fmt"
	"os"
)

// How to use a maps as set, uniq elements
func main() {
	seen := make(map[string]bool) // Resembles set

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprint(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
