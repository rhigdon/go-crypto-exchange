package exchange

import (
	"encoding/json"
	"errors"
	"log"
)

// GetAccountAPI is a helper function to get parsed user account data
func GetAccountAPI(key string, secret string) (Balance, error) {
	api, err := NewUserAPI(key, secret)
	var data Balance
	if err != nil {
		log.Fatalf("Unable to create UserAPI Instance: %s", err.Error())
		return data, err
	}

	resp, err := api.Account()
	if err != nil {
		log.Fatalf("Unable to retrieve Account Balance: %s", err.Error())
		return data, err
	}

	if resp.Code != "0" {
		log.Fatalf("[%s] API Error %s", resp.Code, resp.Message)
		return data, errors.New(resp.Message)
	}

	json.Unmarshal(*resp.Data, &data)
	return data, nil
}

// GetMarketAPI is a helper function to get parsed market data
func GetMarketAPI() (TickerPrice, error) {
	market := NewMarketAPI()
	resp, err := market.TickerPrice()
	var data TickerPrice
	if err != nil {
		log.Fatalf("Unable to retrieve Market Ticker: %s", err.Error())
		return data, err
	}

	if resp.Code != "0" {
		log.Fatalf("[%s] API Error %s", resp.Code, resp.Message)
		return data, errors.New(resp.Message)
	}

	json.Unmarshal(*resp.Data, &data)
	return data, nil
}
