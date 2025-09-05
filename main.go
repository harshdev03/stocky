package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <STOCK_SYMBOL>")
		return
	}

	stockSymbol := os.Args[1]
	url := fmt.Sprintf("https://indian-stock-exchange-api1.p.rapidapi.com/stock_price/?symbol=%s", stockSymbol)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-rapidapi-key", "34659b738bmsheb9bb0c92c2808dp1f71f2jsn082f88e3e848")
	req.Header.Add("x-rapidapi-host", "indian-stock-exchange-api1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var stock Stock
	json.Unmarshal(body, &stock)

	fmt.Printf("Stock: %s\n", stock.Symbol)
	fmt.Printf("Current Price: %.2f\n", stock.LastTradingPrice)
	fmt.Printf("Open: %.2f\n", stock.Open)
	fmt.Printf("Day High: %.2f\n", stock.DayHigh)
	fmt.Printf("Day Low: %.2f\n", stock.DayLow)
	fmt.Printf("Previous Close: %.2f\n", stock.PreviousClose)
	fmt.Printf("Day Change: %.2f (%.2f%%)\n", stock.DayChange, stock.DayChangePercent)
	fmt.Printf("52W High: %.2f\n", stock.HighPriceRange)
	fmt.Printf("52W Low: %.2f\n", stock.LowPriceRange)
	fmt.Printf("Volume: %d\n", stock.Volume)
	fmt.Printf("Buy Qty: %d\n", stock.TotalBuyQty)
	fmt.Printf("Sell Qty: %d\n", stock.TotalSellQty)
}
