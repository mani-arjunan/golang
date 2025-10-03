package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
)

var API_KEY = os.Getenv("W_API_KEY")

type Coord struct {
	Lon float32 `json:"lon"`
	Lat float32 `json:"lat"`
}

type Main struct {
	Temperature float32 `json:"temp"`
}

type WeatherResponse struct {
	Main  Main  `json:"main"`
	Coord Coord `json:"coord"`
}

func makeApiCall(ch chan<- string, wg *sync.WaitGroup, city string) {
	var data WeatherResponse

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, API_KEY)
	resp, err := http.Get(url)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		return
	}

	ch <- fmt.Sprintf("Temperature for %s: %f \n", city, data.Main.Temperature)
	fmt.Printf("City %s sent data via channel \n", city)

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)
	cities := []string{"Chennai", "Mumbai", "Nagapattinam", "Hyderabad", "Karur", "Delhi"}

	for _, city := range cities {
		wg.Add(1)
		go makeApiCall(ch, &wg, city)
	}

	go func() {
		wg.Wait()
		close(ch)
		fmt.Println("channel Closed")
	}()

	for data := range ch {
		fmt.Println("priting in main", data)
	}
	fmt.Println("main Closed")

	fmt.Println("Main Program end")
}
