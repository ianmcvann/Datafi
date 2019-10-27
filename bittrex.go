package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
type BittrexResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  struct {
		Bid  float64 `json:"Bid"`
		Ask  float64 `json:"Ask"`
		Last float64 `json:"Last"`
	} `json:"result"`
}
func makeBittrexTickerResponse(ticker string, data string) BittrexResponse {
	var response BittrexResponse
	dataByte := []byte(data)
	err := json.Unmarshal(dataByte, &response)
	if err != nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
	}
	return response
}
func getBittrexPairInfo(ticker string) (string, string) {
	firstTicker := ticker[0:3]
	secondTicker := ticker[len(ticker)-3:]
	request := fmt.Sprintf("https://api.bittrex.com/api/v1.1/public/getticker?market=%v-%v", secondTicker, firstTicker)
	response, err := http.Get(request)
	if err != nil {
		return ticker, fmt.Sprintf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		return ticker, string(data)
	}
}
