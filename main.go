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
	Name   string  `json:"name"`
	Price  float64 `json:"price,string"`
	Open   float64 `json:"open,string"`
	High   float64 `json:"high,string"`
	Low    float64 `json:"low,string"`
	Volume int64   `json:"volume,string"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	apiKey := os.Getenv("RAPIDAPI_KEY")
	if apiKey == "" {
		fmt.Println("API key not found. Please set RAPIDAPI_KEY in .env file.")
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Please provide a stock name: ")
		fmt.Println("Example: go run main.go \"Tata Steel\"")
		return
	}

	stockName := strings.Join(os.Args[1:], " ")
	url := fmt.Sprintf("https://indian-stock-exchange-api2.p.rapidapi.com/stock?name=%s", stockName)

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

	var stock Stock
	err = json.Unmarshal(body, &stock)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		fmt.Println("Response:", string(body))
		return
	}

	fmt.Printf("\n📊 Stock: %s\n", stock.Name)
	fmt.Printf("💵 Price: %.2f\n", stock.Price)
	fmt.Printf("💰 Open: %.2f\n", stock.Open)
	fmt.Printf("📈 High: %.2f\n", stock.High)
	fmt.Printf("📉 Low: %.2f\n", stock.Low)
	fmt.Printf("🔄 Volume: %d\n\n", stock.Volume)
}
