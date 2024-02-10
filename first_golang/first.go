package main

import "fmt"
import "first/nested"
import "rsc.io/quote"

func main() {
  fmt.Println("Hello! this is my first go program")
  // Usage of local module
  nested.Nes()
  // Usage of external module functions
  fmt.Println(quote.Go())
}


