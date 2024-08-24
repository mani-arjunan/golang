package main

import "fmt"

func main() {
	var value = 10
	ptr1 := &value
	ptr2 := &ptr1
  ptr3 := &ptr2

  **ptr2 = 200
	fmt.Printf("Actual Value ===> %v", ***ptr3)

}
