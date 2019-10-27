package main

//Copyright 2019 Ian McVann
import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

const (
	host     = "192.168.0.3"
	port     = 5432
	user     = "USER"
	password = "PASS"
	dbname   = "datafiwrite"
)

type Tracker struct {
	Exchange string
	Ticker   string
}

func (t Tracker) testConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
func (t Tracker) track() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := `CREATE TABLE IF NOT EXISTS "kraken"
	(id serial primary key, dtime text, ticker text, ask real, bid real, volume real, volume_weighted_avg_price real, trade_count integer, open real, high real, low real, close real)`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	for {
		dataset := makeKrakenTickerResponse(getKrakenPairInfo(t.Ticker))
		for _, data := range dataset {
			fmt.Printf("Logging Data for ticker %v\n", data.Ticker)
			dt := time.Now()
			sqlStatement = `INSERT INTO "kraken" VALUES (DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
			_, err = db.Exec(sqlStatement, dt, data.Ticker, data.Ask, data.Bid, data.Volume, data.VolumeWeightedPrice, data.TradeCount, data.Open, data.High, data.Low, data.Close)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Data for %v logged.\n", data.Ticker)
		}
		time.Sleep(time.Hour)
	}
}
