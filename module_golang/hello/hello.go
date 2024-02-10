package main

import "fmt"
import "module_example/welcome"

func main() {
  var message string
  message = welcome.Hello("Manikandan")
  fmt.Println(message)
}
