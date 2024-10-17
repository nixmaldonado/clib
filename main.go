package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/text/message"
	"net/http"
	"os"
)

type CoinGeckoResponse struct {
	Bitcoin struct {
		USD          float64 `json:"usd"`
		USDMarketCap float64 `json:"usd_market_cap"`
	} `json:"bitcoin"`
}

func main() {
	cgResponse, err := getBTCPrice()
	if err != nil {
		fmt.Println("Error fetching the Bitcoin price:", err)
		os.Exit(1)
	}

	p := message.NewPrinter(message.MatchLanguage("en"))
	formattedPrice := p.Sprintf("%.2f", cgResponse.Bitcoin.USD)
	formattedMarketCap := p.Sprintf("%.2f", cgResponse.Bitcoin.USDMarketCap)

	fmt.Printf("Current Bitcoin Price: U$D%s \n", formattedPrice)
	fmt.Printf("Current Bitcoin Market Cap: U$D%s\n", formattedMarketCap)
}

// Function to fetch the current price of Bitcoin
func getBTCPrice() (CoinGeckoResponse, error) {
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd&include_market_cap=true")
	var result CoinGeckoResponse

	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}
