package main

import ("fmt"; "sort")

func main() {
	var names []string

	ages := make(map[string]int)

	ages["Marti"] = 29
	ages["Kremi"] = 25


	for name := range ages {
		names = append(names,name)
	}

	sort.Strings(names) // Sorting keys

	// Printing the map sorted
	for _ ,v := range names{ 
		fmt.Printf("Name : %s , Age : %d\n",v,ages[v])
	}


}