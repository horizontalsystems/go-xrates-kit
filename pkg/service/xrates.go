package service

import "github.com/horizontalsystems/xrates-kit/pkg/handler"
import "github.com/horizontalsystems/xrates-kit/pkg/config"

var conf = config.Get()
var ipfsHandler handler.XRates
var coinPaprika handler.XRates

func init() {
	ipfsHandler = handler.Ipfs{&(conf.Ipfs)}
	coinPaprika = handler.CoinPaprika{&conf.CoinPaprika}
}

// GetLatestXRates gets latest rates of source and target currencies
func GetLatestXRates(dCcy string, fCcy string, exchange string) string {
	data, err := ipfsHandler.GetLatestXRates(dCcy, fCcy, exchange)

	if err != nil {

		// Get data from CoinPaprika
		data, err := coinPaprika.GetLatestXRates(dCcy, fCcy, exchange)

		if err != nil {

		}

		return data
	}

	return data
}

// GetXRates gets rates by date
func GetXRates(dCcy string, fCcy string, exchange string, epochSec int64) string {
	data, err := ipfsHandler.GetXRates(dCcy, fCcy, exchange, &epochSec)

	if err != nil {

	}

	return data
}
