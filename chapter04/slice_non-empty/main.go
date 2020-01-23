package main

import "fmt"

func main() {

	a := []string{"Marti", "", "", "Katya", "", "Kremema"}

	//fmt.Println(non_empty(a))
	fmt.Println(non_empty2(a))

}

func non_empty(strings []string) []string {

	i := 0
	for _, str := range strings {
		if str != "" {
			strings[i] = str
			i++ // Increses only when non-empty element
		}
	}

	return strings[:i] // Returning what is not empty

}

func non_empty2(strings []string) []string {


	return_slice := make([]string, 0)

	fmt.Println(return_slice)

	for _, v := range strings {
		if v != "" {
			return_slice = append(return_slice, v)
		}
	}

	return return_slice
}
