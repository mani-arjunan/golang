package welcome

import (
	"errors"
	"fmt"
	"math/rand"
)

// Basic function with error handling
// func Hello(name string) (string, error) {
//   // Initialize and assigns value in on line, where the type will be determined by the right side value
//   if name == "" {
//     return "", errors.New("Empty name is not allowed!")
//   }
//   // Sprintf doesn't log into console it takes string along with format specifier like %v etc just like c and
//   // returns the output string by replacing those format specifiers with exact values
//   message := fmt.Sprintf("Hi, %v. Welcome!", name)
//   // alternate or tradinational approach for the above line
//   // var message string
//   // message = fmt.Sprintf("Hi, %v. Welcome!", name)
//   // return nil in place of error
//   return message, nil
// }

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name is not allowed!")
	}

	message := fmt.Sprintf(generateRandomWelcomeMsg(), name)
  // For Testing purpose
	// message := fmt.Sprintf(generateRandomWelcomeMsg())

	return message, nil
}

// returns a map that associates 
func Hellos(names []string) (map[string]string, error) {
  // map[KEY_TYPE]VALUE_TYPE
  // map[string]number is similar to Record<string, number>

  // make is similar to malloc in c
  // map are reference types like pointers in c, so the value of messages
  // could be a garbage value or nil in go terms, so when initilizing this will cause run time error 
  // since we cannot initilize in an unknown or nil memory space, so in order to create a virtual memory space
  // make is used
  // make will allocate and initiliazes a virtual memory for the created map
  var messages map[string]string
  messages = make(map[string]string)

  // Go for loop
  // i => index, name => value
  for _, name := range names {
    message, err := Hello(name)
    if err != nil {
      return nil, err
    }

    messages[name] = message
  }

  return messages, nil
}

// Note: its a lower case function, means it can only be used internally, not exported
func generateRandomWelcomeMsg() string {
	// declare an array of strings which can be returned based on random values
	welcomeMessages := []string{
		"Hi! %v welcome",
		"Hello! %v welcome! Nice to see you",
		"Hello! %v Happy to see you",
	}

	// rand is inbuild function similar to Math.random in js, here it takes a number as an argument and it generates random numbers from 0 - number
	return welcomeMessages[rand.Intn(len(welcomeMessages))]
}
