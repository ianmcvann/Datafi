package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)
func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Datafi API Home.")
}
func getTokenLink(w http.ResponseWriter, r *http.Request) {
	ticker := string(mux.Vars(r)["ticker"])
	response, _ := json.MarshalIndent(makeTickerResponse(getPairInfo(ticker)), "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(response))
}
func main() {
	var tracker Tracker
	tracker.Exchange = "kraken"
	tracker.track()
	//router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/", homeLink)
	//router.HandleFunc("/{ticker}", getTokenLink)
	//log.Fatal(http.ListenAndServe(":8080", router))

}
