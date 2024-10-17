package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"golang.org/x/text/message"
	"net/http"
	"os"
)

type CoinGeckoResponse struct {
	Bitcoin struct {
		USD float64 `json:"usd"`
	} `json:"bitcoin"`
}

func main() {
	name := flag.String("name", "World", "your name")

	getBtc := flag.Bool("btc", false, "Fetch the current Bitcoin price in USD")

	flag.Parse()

	fmt.Printf("Hello, %s!\n", *name)

	if *getBtc {
		price, err := getBTCPrice()
		if err != nil {
			fmt.Println("Error fetching the Bitcoin price:", err)
			os.Exit(1)
		}

		p := message.NewPrinter(message.MatchLanguage("en"))
		formattedPrice := p.Sprintf("%.2f", price)

		fmt.Printf("The current price of Bitcoin is: $%s USD\n", formattedPrice)
	}
}

// Function to fetch the current price of Bitcoin
func getBTCPrice() (float64, error) {
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result CoinGeckoResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	return result.Bitcoin.USD, nil
}
