package xrates

import (
	"strings"

	"github.com/horizontalsystems/xrates-kit/pkg/config"
	"github.com/horizontalsystems/xrates-kit/pkg/handler"
)

var conf = config.Get()
var ipfsHandler, coinPaprika handler.XRates

func init() {
	ipfsHandler = handler.Ipfs{&(conf.Ipfs)}
	coinPaprika = handler.CoinPaprika{&conf.CoinPaprika}
}

// GetLatest gets latest rates of source and target currencies
func GetLatest(digCurrency string, fiatCurrency string, exchange string) string {
	data, err := ipfsHandler.GetLatestXRates(digCurrency, fiatCurrency, exchange)

	fiatCurrency = strings.ToUpper(fiatCurrency)
	digCurrency = strings.ToUpper(digCurrency)

	if err != nil {

		// Get data from CoinPaprika
		data, err := coinPaprika.GetLatestXRates(digCurrency, fiatCurrency, exchange)

		if err != nil {

		}

		return data
	}

	return data
}

// Get gets rates by Unix EPOCH date
func Get(digCurrency string, fiatCurrency string, exchange string, epochSec int64) string {

	fiatCurrency = strings.ToUpper(fiatCurrency)
	digCurrency = strings.ToUpper(digCurrency)

	data, err := ipfsHandler.GetXRates(digCurrency, fiatCurrency, exchange, &epochSec)

	if err != nil {

	}

	return data
}
