package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/metarsit/exchange"
)

func main() {
	var symbols string

	flag.StringVar(&symbols, "symbols", "", "Name of a markets with comma delimited")
	flag.Parse()

	data, err := exchange.GetMarketAPI()
	if err != nil {
		os.Exit(1)
	}

	if symbols == "" {
		for ticker, price := range data {
			fmt.Printf("%s: %v\n", strings.ToUpper(ticker), price)
		}
	} else {
		var symbolList = strings.Split(symbols, ",")
		for _, symbol := range symbolList {
			fmt.Printf("%s: %v\n", strings.ToUpper(symbol), data[strings.ToLower(symbol)])
		}
	}
}
