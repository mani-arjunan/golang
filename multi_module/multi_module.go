package main 

import "fmt"
import "golang.org/x/example/hello/reverse"


func main() {
  fmt.Println(reverse.String("Hello"), reverse.Int(24601))
}
