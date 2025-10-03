// This is a dumb example and i know, i just wrote it to understand channels vs mutex usecase
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

var data WeatherResponse

func makeApiCall(mu *sync.Mutex, wg *sync.WaitGroup, city string) {

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, API_KEY)
	resp, err := http.Get(url)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	mu.Lock()
	err = json.NewDecoder(resp.Body).Decode(&data)
	fmt.Printf("%s: %f\n", city, data.Main.Temperature)
	mu.Unlock()

	if err != nil {
		return
	}

	// ch <- fmt.Sprintf("Temperature for %s: %f \n", city, data.Main.Temperature)

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	cities := []string{"Chennai", "Mumbai", "Nagapattinam", "Hyderabad", "Karur", "Delhi"}

	for _, city := range cities {
		wg.Add(1)
		go makeApiCall(&mu, &wg, city)
	}

	wg.Wait()
	fmt.Printf("City %f sent data via channel \n", data.Main.Temperature)

	fmt.Println("Main Program end")
}
