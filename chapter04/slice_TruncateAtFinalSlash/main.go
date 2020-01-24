package main

import (
	"bytes"
	"fmt"
)

type path []byte

func (p *path) TruncateAtFinalSlash() {
	i := bytes.LastIndex(*p, []byte("/"))

	if i >= 0 {
		*p = (*p)[0:i]
	}

}

func main() {
	pathName := path("/usr/marti/docs")
	pathName.TruncateAtFinalSlash() //  Method , it uses the same slice header and underlying array
	fmt.Printf("%s\n", pathName)
}

