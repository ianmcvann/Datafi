package main

type Tracker struct {
	Exchange string
	Ticker string
}

func (t Tracker) track() {
	getPairInfo(t.Ticker)
}