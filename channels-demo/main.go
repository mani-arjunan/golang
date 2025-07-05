package main

import (
	"channels-demo/api"
	"fmt"
	"sync"
)

func main() {
	currencies := [3]string{"BTC", "ETH", "BCH"}
	var wg sync.WaitGroup

	for _, val := range currencies {
		wg.Add(1)
		go callAPI(val, &wg)
	}

	wg.Wait()
}

func callAPI(currency string, wg *sync.WaitGroup) {
	res, err := api.GetRate(currency)

	if err != nil {
		fmt.Println("Err in getting the data")
	}

	fmt.Println(res.Price, "FOR", currency)
	wg.Done()
}

// My version of sync.WaitGroup
// package main
//
// import (
// 	"channels-demo/api"
// 	"fmt"
// )
//
// func main() {
// 	currencies := [3]string{"BTC", "ETH", "BCH"}
// 	count := 0
//
// 	for _, val := range currencies {
// 		count++
// 		go callAPI(val, &count)
// 	}
//
// 	for {
// 		if count == 0 {
// 			break
// 		}
// 	}
// }
//
// func callAPI(currency string, count *int) {
// 	res, err := api.GetRate(currency)
//
// 	if err != nil {
// 		fmt.Println("Err in getting the data")
// 	}
//
// 	fmt.Println(res.Price, "FOR", currency)
// 	(*count)--
// }
