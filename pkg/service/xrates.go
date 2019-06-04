package service

import (
	"strings"

	"github.com/horizontalsystems/xrates-kit/pkg/config"
	"github.com/horizontalsystems/xrates-kit/pkg/handler"
)

type XRatesService struct {
	conf                     *config.MainConfig
	ipfsHandler, coinPaprika handler.XRates
}

// Init x-rates service, init and load configurations for Handlers
func (xratesSrv *XRatesService) Init() {
	xratesSrv.conf = config.Get()
	xratesSrv.ipfsHandler = &handler.Ipfs{&(xratesSrv.conf.Ipfs)}
	xratesSrv.coinPaprika = &handler.CoinPaprika{&xratesSrv.conf.CoinPaprika}
}

// GetLatest gets latest rates of source and target currencies
func (xratesSrv *XRatesService) GetLatest(
	digCurrency string, fiatCurrency string, exchange string) string {

	data, err := xratesSrv.ipfsHandler.GetLatestXRatesAsJSON(digCurrency, fiatCurrency, exchange)

	fiatCurrency = strings.ToUpper(fiatCurrency)
	digCurrency = strings.ToUpper(digCurrency)

	if err != nil {

		// Get data from CoinPaprika
		data, err := xratesSrv.coinPaprika.GetLatestXRatesAsJSON(digCurrency, fiatCurrency, exchange)

		if err != nil {

		}

		return data
	}

	return data
}

// Get method gets rates by Unix EPOCH date
func (xratesSrv *XRatesService) Get(
	digCurrency string, fiatCurrency string, exchange string, epochSec int64) string {

	fiatCurrency = strings.ToUpper(fiatCurrency)
	digCurrency = strings.ToUpper(digCurrency)

	data, err := xratesSrv.ipfsHandler.GetXRatesAsJSON(digCurrency, fiatCurrency, exchange, &epochSec)

	if err != nil {
		// Get data from CoinPaprika
		data, err := xratesSrv.coinPaprika.GetXRatesAsJSON(digCurrency, fiatCurrency, exchange, &epochSec)

		if err != nil {

		}

		return data
	}

	return data
}
