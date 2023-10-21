package main

import (
	"fmt"
	"sync"

	"frontendmasters.com/go/crypto/api"
)

func main() {
	currencies := []string { "BTC", "ETH", "SOL"}
	
	// waitgroup - waits for goroutines to finish
	var wg sync.WaitGroup

for _, currency := range currencies {
	// creating a lamba function
	wg.Add(1)
	go func (currencyCode string) {
		// passing currency as a argument so it is a copy and not a reference
		getCurrencyData(currencyCode)
		wg.Done()
	}(currency)
}
wg.Wait()
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)

	if err == nil {
		fmt.Printf("The rate for %v is %.2f\n", rate.Currency, rate.Price)
	}
}