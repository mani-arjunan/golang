package api

import (
	"channels-demo/datatypes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Rate struct {
	Currency string `json:"pair"`
	Price float64 `json:"last"`
}

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	var finalValue datatypes.Rate
	res, err := http.Get(fmt.Sprintf(apiUrl, strings.ToUpper(currency)))

	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)

		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(bodyBytes, &finalValue)

		if err != nil {
			fmt.Println("ERRRR HERE")
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("Errr: %v", res.StatusCode)
	}

	return &finalValue, err
}
