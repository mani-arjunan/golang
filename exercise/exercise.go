package exercise

import "fmt"

type Item struct {
  Name string
  Type string
}

type Player struct {
  Name string
  Inventory []Item
}

func (p *Player) PickupItem(newItem Item) {
  p.Inventory = append(p.Inventory, newItem)
}

func (p *Player) DropItem(itemName string) {
  tempArr := []Item{}

  for _, value := range p.Inventory {
    if itemName == value.Name {
      continue
    } else {
      tempArr = append(tempArr, value)
    }
  }
  
  p.Inventory = tempArr
}

func functionExample (a int) (c int, b string) {
  a = 1 + 1
  b = "Helo"

  return
}

func incrementAge(age *int) {
  *age = *age + 1
}

// init function will be called at start, there can be multiple init functions
func init() {
  fmt.Println("Hello from exercise")
  var str1 int
  var str2 string

  str1, str2 = functionExample()

  fmt.Println(str1, str2)
  age := 12

  incrementAge(&age)

  fmt.Println(age)
}




