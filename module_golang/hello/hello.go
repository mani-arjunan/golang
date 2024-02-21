package main

import "fmt"
import "log"
import "module_example/welcome"

func main() {
  // Setting predefined prefix for logger
  log.SetPrefix("Welcome: ")

  // Module 1
  // var message string
  // var error error 
  // message, error = welcome.Hello("Manikandan")
  // // alternate for line above
  // // message, error := welcome.Hello("")
  // if error != nil {
  //   log.Fatal(error)
  // }

  // Module 2
  var names []string = []string{"Manikandan", "Hema", "test"}
  var messages map[string]string
  var error error

  messages, error = welcome.Hellos(names)

  if error != nil {
    log.Fatal(error)
  }
  
  // To Print the entire map
  // fmt.Println(messages)
  fmt.Println(messages["Manikandan"])
  fmt.Println(messages["Hema"])
  fmt.Println(messages["test"])
}
