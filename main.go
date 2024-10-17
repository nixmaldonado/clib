package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/text/message"
	"net/http"
	"os"
	"time"
)

type CoinGeckoResponse struct {
	Bitcoin struct {
		USD           float64 `json:"usd"`
		USDMarketCap  float64 `json:"usd_market_cap"`
		LastUpdatedAt int64   `json:"last_updated_at"`
	} `json:"bitcoin"`
}

func main() {
	info, err := getBTCInfo()
	if err != nil {
		fmt.Println("Error fetching the Bitcoin price:", err)
		os.Exit(1)
	}

	p := message.NewPrinter(message.MatchLanguage("en"))
	formattedPrice := p.Sprintf("%.2f", info.Bitcoin.USD)
	formattedMarketCap := formatLargeNumber(info.Bitcoin.USDMarketCap)
	t := time.Unix(info.Bitcoin.LastUpdatedAt, 0)
	localTime := t.Local().Format("2006-01-02 15:04:05")

	fmt.Printf("Bitcoin Price: U$D%s @ %s \n", formattedPrice, localTime)
	fmt.Printf("Market Cap: U$D%s\n", formattedMarketCap)
}

// Function to fetch the current price of Bitcoin
func getBTCInfo() (CoinGeckoResponse, error) {
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd&include_market_cap=true&include_last_updated_at=true")
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

func formatLargeNumber(num float64) string {
	var suffix string
	var value float64

	if num >= 1e12 {
		value = num / 1e12
		suffix = "Trillion"
	} else if num >= 1e9 {
		value = num / 1e9
		suffix = "Billion"
	} else {
		return fmt.Sprintf("%.2f", num)
	}

	// Format with 2 decimal places
	return fmt.Sprintf("%.2f %s", value, suffix)
}
