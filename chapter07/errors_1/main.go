package main

import (
	"errors"
	"fmt"
)

type SyntaxError struct {
	Line int
	Col  int
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("%d:%d: syntax error", e.Line, e.Col)
}

func main() {

	err := errors.New("This is an ")

}
