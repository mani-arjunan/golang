package welcome

import "testing"
import "regexp"


// Test Hello
// Test function should starts with Test
// also test function will take testing pointer as parameter
func TestHello(t *testing.T) {
  var name string
  name = "Manikandan"
  want := regexp.MustCompile(`\b`+name+`\b`)

  var msg string
  var error error

  msg, error = Hello("Manikandan")

  if !want.MatchString(msg) || error != nil {
    t.Fatalf(`Hello("Manikandan" = %q, %v, want match for %#q, nil`, msg, error, want)
  }
}


//Testing empty name in Hello

func TestEmptyNameHello(t *testing.T) {
  var msg string
  var err error
  msg, err = Hello("")

  if msg != "" || err == nil {
    t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
  }
}
