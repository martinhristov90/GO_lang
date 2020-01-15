package main

import ("fmt")


func foo(a int) {
  if (a > 0) {
      fmt.Println("Hi")
  } else {
      // even if never go to else statement, error is thrown, cause  go is static-typed, in Python would pass
      fmt.Println("3" + 5)
  }
}

func main() {
  foo(2)
}