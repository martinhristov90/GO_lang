package main

import (
	"fmt"
	"os"
)

//var i,j int = 1,2
var f,err = os.Open("text.txt") // f,err :=os.Open("text.txt")

func main(){
	//fmt.Println("Printing a variable:",i,j)
	fmt.Println("Printing a variable:",f)
}