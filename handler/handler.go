package handler

import (
	"github.com/horizontalsystems/xrates-kit/models"
)

//----------------------------
const TIMEOUT_GLOBAL = 6 
const TIMEOUT_IPFS = 6 
const TIMEOUT_COINPAPRIKA = 6 
//----------------------------

//typeXRates interface
type XRates interface {

	// GetLatestXRates gets latest rates of source and target currencies
	GetLatestXRates(cCoin string, fCcy string, exchange string) (*models.LatestXRate, error)

	// Fetches rates by Unix epoch time
	GetHistoricalXRates(cCoin string, fCcy string, exchange string, epochSec *int64) (
		*models.HistoricalXRate, error)
}
