package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Datafi API Home.")
}
func getTokenLink(w http.ResponseWriter, r *http.Request) {
	exchange := string(mux.Vars(r)["exchange"])
	ticker := string(mux.Vars(r)["ticker"])
	var response []byte
	if exchange == "kraken" {
		response, _ = json.MarshalIndent(makeKrakenTickerResponse(getKrakenPairInfo(ticker)), "", "  ")
	} else if exchange == "cex" {
		response, _ = json.MarshalIndent(makeCexTickerResponse(getCexPairInfo(ticker)), "", "  ")
	} else if exchange == "bittrex" {
		response, _ = json.MarshalIndent(makeBittrexTickerResponse(getBittrexPairInfo(ticker)), "", "  ")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(response))
}
func main() {
	var tracker Tracker
	tracker.Exchange = "kraken"
	tracker.Ticker = "BCHEUR,BCHUSD,BCHXBT,DASHEUR,DASHUSD,DASHXBT,EOSETH,EOSXBT,GNOETH,GNOXBT,USDTZUSD,XETCXETH,XETCXXBT,XETCZEUR,XETCZUSD,XETHXXBT,XETHZGBP,XETHZJPY,XETHZUSD,XICNXETH,XICNXXBT,XLTCXXBT,XLTCZEUR,XLTCZUSD,XMLNXETH,XMLNXXBT,XREPXETH,XREPXXBT,XREPZEUR,XXBTZCAD,XXBTZEUR,XXBTZGBP,XXBTZJPY,XXBTZUSD,XXDGXXBT,XXLMXXBT,XXMRXXBT,XXMRZEUR,XXMRZUSD,XXRPXXBT,XXRPZEUR,XXRPZUSD,XZECXXBT,XZECZEUR,XZECZUSD"
	go tracker.track()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/{exchange}/{ticker}", getTokenLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}
