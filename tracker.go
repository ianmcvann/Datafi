package main
//Copyright 2019 Ian McVann
import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "192.168.0.3"
	port     = 5432
	user     = "***REMOVED***"
	password = "***REMOVED***"
	dbname   = "datafiwrite"
)
type Tracker struct {
	Exchange string
	Ticker string
}
func (t Tracker) testConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("***REMOVED***", psqlInfo)
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
	getPairInfo(t.Ticker)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("***REMOVED***", psqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := `CREATE TABLE IF NOT EXISTS "kraken"
	(index integer, dtime text, ticker text, ask real, bid real, close real, volume real, volume_weighted_avg_price real, trade_count integer, low real, high real, open real)`
	_, err = db.Exec(sqlStatement, t.Exchange)
	if err != nil {
		panic(err)
	}
}
