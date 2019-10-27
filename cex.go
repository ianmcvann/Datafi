package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CexResponse struct {
	Timestamp             string  `json:"timestamp"`
	Low                   string  `json:"low"`
	High                  string  `json:"high"`
	Last                  string  `json:"last"`
	Volume                string  `json:"volume"`
	Volume30D             string  `json:"volume30d"`
	Bid                   float64 `json:"bid"`
	Ask                   float64 `json:"ask"`
	PriceChange           string  `json:"priceChange"`
	PriceChangePercentage string  `json:"priceChangePercentage"`
	Pair                  string  `json:"pair"`
}

func makeCexTickerResponse(ticker string, data string) CexResponse {
	var response CexResponse
	dataByte := []byte(data)
	err := json.Unmarshal(dataByte, &response)
	if err != nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
	}
	return response
}
func getCexPairInfo(ticker string) (string, string) {
	firstTicker := ticker[0:3]
	secondTicker := ticker[len(ticker)-3:]
	request := fmt.Sprintf("https://cex.io/api/ticker/%v/%v", firstTicker, secondTicker)
	response, err := http.Get(request)
	if err != nil {
		return ticker, fmt.Sprintf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		return ticker, string(data)
	}
}
