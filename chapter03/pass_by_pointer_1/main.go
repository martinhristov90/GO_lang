package main

import "fmt"

type CPU struct {
	freq float32
	model string
}

func set_def_model(proc *CPU) {
	proc.model = "generic_intel"
}


func main() {

	myproc := CPU {
		freq: 3.22 ,
		model: "i5-7600K",
	}

	// Passing by pointer, so set_def_model will modify the myproc struct,proc and myproc point to the SAME data struct
	set_def_model(&myproc)

	fmt.Println(myproc.model)

}