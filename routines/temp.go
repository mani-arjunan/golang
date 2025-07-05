package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0
var mutex sync.Mutex

func worker(id int) {
  mutex.Lock()
  count++
	fmt.Println("Worker", id, "updated count to", count)
  mutex.Unlock()
}

func main() {
  for i := 0; i < 10; i++ {
    go worker(i)
  }

  time.Sleep(4 * time.Second)
  mutex.Lock()
  fmt.Println("Value of count is ", count)
  mutex.Unlock()
}
