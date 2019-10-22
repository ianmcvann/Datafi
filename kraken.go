package main

import (
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"net/http"
)
type Response struct {
	Ticker string
	Ask string
	Bid string
	Volume string
	VolumeWeightedPrice string
	TradeCount string
	Open string
	High string
	Low string
	Close string
}
func makeTickerResponse(ticker string, data string) Response {
	var response Response
	dataByte := []byte(data)
	paths := [][]string{
		[]string{"result", ticker, "a", "[0]"},
		[]string{"result", ticker, "b", "[0]"},
		[]string{"result", ticker, "c", "[0]"},
		[]string{"result", ticker, "v", "[0]"},
		[]string{"result", ticker, "p", "[0]"},
		[]string{"result", ticker, "t", "[0]"},
		[]string{"result", ticker, "l", "[0]"},
		[]string{"result", ticker, "h", "[0]"},
		[]string{"result", ticker, "o", "[0]"},
	}
	jsonparser.EachKey(dataByte, func(idx int, value []byte, vt jsonparser.ValueType, err error){
		switch idx {
		case 0:
			response.Ask = string(value)
		case 1:
			response.Bid = string(value)
		case 2:
			response.Close = string(value)
		case 3:
			response.Volume = string(value)
		case 4:
			response.VolumeWeightedPrice = string(value)
		case 5:
			response.TradeCount = string(value)
		case 6:
			response.Low = string(value)
		case 7:
			response.High = string(value)
		case 8:
			response.Open = string(value)
		}
	}, paths...)
	response.Ticker = ticker
	return response
}
func getPairInfo(ticker string) (string, string) {
	request := fmt.Sprintf("https://api.kraken.com/0/public/Ticker?pair=%v", ticker)
	response, err := http.Get(request)
	if err != nil {
		return ticker, fmt.Sprintf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		return ticker, string(data)
	}
}
func main() {
	fmt.Println("Starting the application...")
	ticker, response := getPairInfo("BCHUSD")
	r := makeTickerResponse(ticker, response)
	fmt.Printf("Response: %v \n", r)
	fmt.Println("Terminating the application...")
}