package welcome 

import "fmt"

func Hello(name string) string {
  // Initialize and assigns value in on line, where the type will be determined by the right side value
  message := fmt.Sprintf("Hi, %v. Welcome!", name)
  // alternate or tradinational approach for the above line
  // var message string
  // message = fmt.Sprintf("Hi, %v. Welcome!", name)
  return message
}

