package main

import(
	"fmt"
)

func main(){

	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n",x+0/x) // causes devision by zero, panicking
	defer fmt.Printf("defer %d\n",x)
	f(x-1)
}

//f(3)
//f(2)
//f(1)
//defer 1
//defer 2
//defer 3
//panic: runtime error: integer divide by zero when f(0)