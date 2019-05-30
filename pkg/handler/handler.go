package handler

// XRates interface
type XRates interface {

	// GetLatestXRates gets latest rates of source and target currencies
	GetLatestXRates(cCoin string, fCcy string, exchange string) (string, error)

	// Fetches rates by Unix epoch time
	GetXRates(cCoin string, fCcy string, exchange string, epochSec *int64) (string, error)
}

// LatestXRatesData is a container for latest exchange rates
type LatestXRatesData struct {
	Currency string `json:"currency"`
	TimeStr  string `json:"time_str"`
	Time     int64  `json:"time"`

	Rates map[string]string `json:"rates"`
}

// XRatesData is a container for exchange rates
type XRatesData struct {
	Currency string `json:"currency"`
	Rate     string `json:"rate"`
	Time     int64  `json:"time"`
}
