package datatypes

type Rate struct {
	Currency string `json:"pair"`
	Price string `json:"last"`
}

type GetPrices interface {
  GetRate() (string, error)
}

