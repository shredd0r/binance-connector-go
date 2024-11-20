package main

import (
	"context"
	"fmt"

	binance_connector "github.com/shredd0r/binance-connector-go"
)

func main() {
	TickerPrice()
}

func TickerPrice() {
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient("", "", baseURL)

	// TickerPrice
	tickerPrice, err := client.NewTickerPriceService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(tickerPrice))
}
