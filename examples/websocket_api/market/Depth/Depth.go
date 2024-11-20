package main

import (
	"context"
	"fmt"
	"log"

	binance_connector "github.com/shredd0r/binance-connector-go"
)

func main() {
	DepthExample()
}

func DepthExample() {
	client := binance_connector.NewWebsocketAPIClient("", "", "wss://ws-api.testnet.binance.vision/ws-api/v3")
	err := client.Connect()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	defer client.Close()

	response, err := client.NewDepthService().Symbol("").Do(context.Background())
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Println(binance_connector.PrettyPrint(response))

	client.WaitForCloseSignal()
}
