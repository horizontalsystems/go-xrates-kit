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

func createTable(db *sql.DB) {

	log.Print("Creating table")
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS xrates " +
		"( id INTEGER PRIMARY KEY AUTOINCREMENT, " +
		"  timestamp  INTEGER NOT NULL,  " +
		"  source_ccy VARCHAR(6) NOT NULL,  " +
		"  target_ccy VARCHAR(6) NOT NULL,  " +
		"  exchange   VARCHAR(10) ,  " +
		"  value      TEXT NOT NULL)")

	statement.Exec()

}

func (srv *CacheService) Init() {

	log.Print("Opening database")

	dbase, err := sql.Open("sqlite3", srv.DBPath+string(filepath.Separator)+"xrates.db")

	if err != nil {

	}

	srv.db = dbase

	createTable(dbase)
}

func (srv *CacheService) Get(sourceCurrency string, targetCurrency string,
	exchange string, epochSec int64) *models.HistoricalXRate {

	stmt, err := srv.db.Prepare("SELECT value FROM xrates WHERE " +
		"source_ccy = ? AND target_ccy = ? AND timestamp = ?")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var jsonValue string
	err = stmt.QueryRow(sourceCurrency, targetCurrency, epochSec).Scan(&jsonValue)
	if err != nil {
		//log.Fatal(err)
		return nil
	}
	hRate := models.HistoricalXRate{}

	return hRate.FromJsonString(jsonValue)
}

func (srv *CacheService) Set(sourceCurrency string, targetCurrency string,
	exchange string, epochSec int64, xrate *models.HistoricalXRate) {

	log.Println("Begin setting data")

	stmt, err := srv.db.Prepare("INSERT INTO xrates(source_ccy,target_ccy,timestamp,value) VALUES(?,?,?,?)")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(sourceCurrency, targetCurrency, epochSec, xrate.ToJsonString())

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Data saved to Cache")
}
