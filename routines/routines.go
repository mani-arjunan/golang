// package main
//
// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )
//
// var count = 0
//
// func runInThread() {
// 	count++
// 	fmt.Println(count)
// 	time.Sleep(5 * time.Second)
// 	fmt.Println("Completed", count)
// }
//
// func main() {
//   fmt.Println("Maximum threads", runtime.GOMAXPROCS(0))
// 	fmt.Println("Start: Goroutines =", runtime.NumGoroutine())
// 	for i := 0; i < 10; i++ {
// 		go runInThread()
// 	}
//
// 	time.Sleep(100 * time.Millisecond)
// 	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())
// 	time.Sleep(6 * time.Second)
// 	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())
// }
