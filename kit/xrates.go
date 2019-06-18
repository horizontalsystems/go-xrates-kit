package kit

import (
	"github.com/horizontalsystems/go-xrates-kit/models"
	"log"
	"strings"

	"github.com/horizontalsystems/go-xrates-kit/cache"
	"github.com/horizontalsystems/go-xrates-kit/config"
	"github.com/horizontalsystems/go-xrates-kit/handler"
)

type XRatesKit struct {
	conf                     *config.MainConfig
	ipfsHandler, coinPaprika handler.XRates
	cacheService             *cache.CacheService
}

type XRate models.XRate

// Init x-rates service, init and load configurations for Handlers
func (xratesKit *XRatesKit) Init(dbPath string) {

	xratesKit.conf = config.Load()
	xratesKit.ipfsHandler = &handler.Ipfs{&(xratesKit.conf.Ipfs)}
	xratesKit.coinPaprika = &handler.CoinPaprika{&xratesKit.conf.CoinPaprika}

	xratesKit.cacheService = &cache.CacheService{DBPath: dbPath}
	xratesKit.cacheService.Init()

}

func (xratesKit *XRatesKit) GetHistorical(currencyCode string, coinCode string, timestamp int64) *XRate {
	currencyCode = strings.ToUpper(currencyCode)
	coinCode = strings.ToUpper(coinCode)

	data, err := xratesKit.ipfsHandler.GetHistoricalXRates(currencyCode, coinCode, &timestamp)

	if err != nil {
		// Get data from CoinPaprika
		data, err = xratesKit.coinPaprika.GetHistoricalXRates(currencyCode, coinCode, &timestamp)

		if err != nil {

		}
	}

	//---------- Cache Data -------------------
	xratesKit.cacheService.SetHistorical(data)

	log.Println("After set")
	//-----------------------------------------

	return (*XRate)(data)
}

func (xratesKit *XRatesKit) GetHistoricalCached(currencyCode string, coinCode string, timestamp int64) *XRate {
	currencyCode = strings.ToUpper(currencyCode)
	coinCode = strings.ToUpper(coinCode)

	return (*XRate)(xratesKit.cacheService.GetHistorical(coinCode, currencyCode, timestamp))
}

//GetLatest gets latest rates of source and target currencies
func (xratesKit *XRatesKit) GetLatest(currencyCode string, coinCodes []string) *[]XRate {

	currencyCode = strings.ToUpper(currencyCode)
	//coinCode = strings.ToUpper(coinCode)

	data, err := xratesKit.ipfsHandler.GetLatestXRates(currencyCode, &coinCodes)

	if err != nil {

		// Get data from CoinPaprika
		data, err = xratesKit.coinPaprika.GetLatestXRates(currencyCode, &coinCodes)

		if err != nil {
			//TODO
		}
	}
	//---------- Cache Data -------------------

	xratesKit.cacheService.SetLatest(data)

	//-----------------------------------------

	log.Println(data)

	var result = make([]XRate, len(*data))
	for i, xrate := range *data {
		result[i] = XRate(xrate)
	}
	return &result
}

//GetLatest gets latest rates of source and target currencies
func (xratesKit *XRatesKit) GetLatestCached(currencyCode string, coinCodes []string) *[]XRate {

	currencyCode = strings.ToUpper(currencyCode)
	result := make([]XRate, len(coinCodes))

	for i, coinCode := range coinCodes {
		coinCode = strings.ToUpper(coinCode)
		xrate := xratesKit.cacheService.GetLatest(coinCode, currencyCode)

		result[i] = XRate(*xrate)

	}
	return &result
}

//
////---------------------------Subscription------------------------------------
//
//type LatestRateListener interface {
//	 OnUpdate(rates []XRate)
//}
//
//func (xratesKit *XRatesKit)subscribeToLatestRates(currencyCode string, coinCode []string, listener LatestRateListener) {
//
//	// every 3minutes fetch latest rate
//	// save to cache
//	// call listener.OnUpdate
//
//}
