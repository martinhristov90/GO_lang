package main

import ( 
	"fmt"
	"strings"
)

func basename(s string) string {
	// Getting the last index of slash
	slash := strings.LastIndex(s,"/")
	// Removing everything before it
	s = s[slash+1:]
	// if it is like ".terraform-provider-sops_v1"
	if dot := strings.LastIndex(s,"."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {

	fmt.Println(basename("/Users/marti/Downloads/terraform-provider-sops_v1"))

}