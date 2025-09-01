package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Stock struct {
	Symbol           string  `json:"symbol"`
	Open             float64 `json:"open"`
	DayHigh          float64 `json:"day_high"`
	DayLow           float64 `json:"day_low"`
	PreviousClose    float64 `json:"previous_close"`
	LastTradingPrice float64 `json:"last_trading_price"`
	LowPriceRange    float64 `json:"lowPriceRange"`
	HighPriceRange   float64 `json:"highPriceRange"`
	Volume           int64   `json:"volume"`
	DayChange        float64 `json:"day_change"`
	DayChangePercent float64 `json:"day_change_percent"`
	TotalBuyQty      int64   `json:"totalBuyQty"`
	TotalSellQty     int64   `json:"totalSellQty"`
}

func main() {
	// load .env
	_ = godotenv.Load()

	apiKey := os.Getenv("RAPIDAPI_KEY")
	if apiKey == "" {
		fmt.Println("API key not found. Please set RAPIDAPI_KEY in .env file.")
		return
	}

	// check args
	if len(os.Args) < 2 {
		fmt.Println("Please provide a stock name: ")
		fmt.Println("Example: go run main.go \"TATAMOTORS\"")
		return
	}

	// build url
	stockName := strings.Join(os.Args[1:], " ")
	url := fmt.Sprintf("https://indian-stock-exchange-api2.p.rapidapi.com/stock?name=%s", stockName)

	// request
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-rapidapi-key", apiKey)
	req.Header.Add("x-rapidapi-host", "indian-stock-exchange-api2.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	// debug: print raw response (uncomment if needed)
	// fmt.Println("Raw Response:", string(body))

	var stock Stock
	err = json.Unmarshal(body, &stock)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		fmt.Println("Response:", string(body))
		return
	}

	// show clean output
	fmt.Printf("Symbol: %s\n", stock.Symbol)
	fmt.Printf("Open: %.2f\n", stock.Open)
	fmt.Printf("High: %.2f\n", stock.DayHigh)
	fmt.Printf("Low: %.2f\n", stock.DayLow)
	fmt.Printf("Previous Close: %.2f\n", stock.PreviousClose)
	fmt.Printf("Last Price: %.2f\n", stock.LastTradingPrice)
	fmt.Printf("52W Low: %.2f\n", stock.LowPriceRange)
	fmt.Printf("52W High: %.2f\n", stock.HighPriceRange)
	fmt.Printf("Volume: %d\n", stock.Volume)
	fmt.Printf("Change: %.2f\n", stock.DayChange)
	fmt.Printf("Change %%: %.2f\n", stock.DayChangePercent)
	fmt.Printf("Buy Qty: %d\n", stock.TotalBuyQty)
	fmt.Printf("Sell Qty: %d\n", stock.TotalSellQty)
}
