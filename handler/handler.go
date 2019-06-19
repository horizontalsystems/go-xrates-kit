package handler

import (
	"github.com/horizontalsystems/go-xrates-kit/models"
)

//----------------------------
const TIMEOUT_GLOBAL = 6
const TIMEOUT_IPFS = 6
const TIMEOUT_COINPAPRIKA = 6

//----------------------------

//typeXRates interface
type XRates interface {

	// GetLatestXRates gets latest rates of source and target currencies
	GetLatestXRates(currencyCode string, coinCodes []string) ([]models.XRate, error)

	// Fetches rates by Unix epoch time
	GetHistoricalXRates(currencyCode string, coinCode string, timestamp *int64) (*models.XRate, error)
}
