package models

// a container for exchange rates
type XRate struct {
	CoinCode     string
	CurrencyCode string
	Timestamp    int64
	Rate         string
}
