//“Write an in-place function to eliminate adjacent duplicates in a []string slice.”
// Hint - remove duplicates  return append(s[:index], s[index+1:]...)



package main

import "fmt"

func main() {

	a := []string{"Marti", "Martin", "Katya", "Katya", "Kremi"}
	
	fmt.Println(squash_dup(a))

}

func squash_dup(string_slice []string) []string {
	// Source https://codereview.stackexchange.com/questions/220273/remove-adjacent-duplicates-in-golang
	for i := 0 ; i < len(string_slice) - 1 ; i++ {
		for j := i + 1; j < len(string_slice) ; j++ {
			if string_slice[i] != string_slice[j] {
				break
			} else {
				string_slice = append(string_slice[:i],string_slice[i+1:]...)
			}

		}

	}
	return string_slice
}