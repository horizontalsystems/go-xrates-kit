package cache

import (
	"database/sql"
	"github.com/horizontalsystems/go-xrates-kit/models"
	"log"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type CacheService struct {
	DBPath string
	db     *sql.DB
}

func createHistoricalRatesTable(db *sql.DB) {

	log.Print("Creating historical rates table")

	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS xrates_historical " +
		"( coin_code VARCHAR(6) NOT NULL,  " +
		"  currency_code VARCHAR(6) NOT NULL,  " +
		"  timestamp  INTEGER NOT NULL,  " +
		"  rate      VARCHAR NOT NULL," +
		"  exchange   VARCHAR(10),  " +
		"  PRIMARY KEY(coin_code, currency_code, timestamp)" +
		")")

	statement.Exec()

}

func createLatestRatesTable(db *sql.DB) {

	log.Print("Creating latest rates table")

	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS xrates_latest " +
		"( coin_code VARCHAR(6) NOT NULL,  " +
		"  currency_code VARCHAR(6) NOT NULL,  " +
		"  timestamp  INTEGER NOT NULL,  " +
		"  rate      VARCHAR NOT NULL," +
		"  exchange   VARCHAR(10),  " +
		"  PRIMARY KEY(coin_code, currency_code)" +
		")")

	statement.Exec()

}

func (srv *CacheService) Init() {

	log.Print("Opening database")

	dbase, err := sql.Open("sqlite3", srv.DBPath+string(filepath.Separator)+"xrates.db")

	if err != nil {

	}

	srv.db = dbase

	createHistoricalRatesTable(dbase)
	createLatestRatesTable(dbase)
}

func (srv *CacheService) GetHistorical(coinCode string, currencyCode string, timestamp int64) *models.XRate {

	stmt, err := srv.db.Prepare("SELECT rate FROM xrates_historical WHERE " +
		"coin_code = ? AND currency_code = ? AND timestamp = ?")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var rate string
	err = stmt.QueryRow(coinCode, currencyCode, timestamp).Scan(&rate)
	if err != nil {
		//log.Fatal(err)
		return nil
	}
	return &models.XRate{coinCode, currencyCode, timestamp, rate}
}

func (srv *CacheService) SetHistorical(xRate *models.XRate) {

	log.Println("Begin setting data", xRate)

	stmt, err := srv.db.Prepare("INSERT OR REPLACE INTO xrates_historical(coin_code, currency_code, timestamp, rate) VALUES(?,?,?,?)")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(xRate.CoinCode, xRate.CurrencyCode, xRate.Timestamp, xRate.Rate)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Data saved to Cache")
}

func (srv *CacheService) GetLatest(coinCode string, currencyCode string) *models.XRate {

	stmt, err := srv.db.Prepare("SELECT rate, timestamp FROM xrates_latest WHERE " +
		"coin_code = ? AND currency_code = ?")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var rate string
	var timestamp int64

	err = stmt.QueryRow(coinCode, currencyCode).Scan(&rate, &timestamp)
	if err != nil {
		//log.Fatal(err)
		return nil
	}
	log.Println("Got latest rate from cache ", rate, timestamp)
	return &models.XRate{coinCode, currencyCode, timestamp, rate}
}

func (srv *CacheService) SetLatest(xRates *[]models.XRate) {

	log.Println("Begin setting latest rate ", xRates)

	stmt, err := srv.db.Prepare("INSERT OR REPLACE INTO xrates_latest(coin_code, currency_code, timestamp, rate) VALUES(?,?,?,?)")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, xRate := range *xRates {
		_, err = stmt.Exec(xRate.CoinCode, xRate.CurrencyCode, xRate.Timestamp, xRate.Rate)
	}

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Data saved to Cache")
}
