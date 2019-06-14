package kit

import (
	"github.com/horizontalsystems/xrates-kit/models"
	"log"
	"strings"

	"github.com/horizontalsystems/xrates-kit/cache"
	"github.com/horizontalsystems/xrates-kit/config"
	"github.com/horizontalsystems/xrates-kit/handler"
)

type XRatesKit struct {
	conf                     *config.MainConfig
	ipfsHandler, coinPaprika handler.XRates
	cacheService             *cache.CacheService
}

type  LatestXRate models.LatestXRate
type  HistoricalXRate models.HistoricalXRate

// Init x-rates service, init and load configurations for Handlers
func (xratesKit *XRatesKit) Init(dbPath string) {

	xratesKit.conf = config.Load()
	xratesKit.ipfsHandler = &handler.Ipfs{&(xratesKit.conf.Ipfs)}
	xratesKit.coinPaprika = &handler.CoinPaprika{&xratesKit.conf.CoinPaprika}

	xratesKit.cacheService = &cache.CacheService{DBPath: dbPath}
	xratesKit.cacheService.Init()

}

// GetLatest gets latest rates of source and target currencies
func (xratesKit *XRatesKit) GetLatest(
	coinCode string, currencyCode string, exchange string) *LatestXRate {

	currencyCode = strings.ToUpper(currencyCode)
	coinCode = strings.ToUpper(coinCode)

	data, err := xratesKit.ipfsHandler.GetLatestXRates(coinCode, currencyCode, exchange)

	if err != nil {

		// Get data from CoinPaprika
		data, err = xratesKit.coinPaprika.GetLatestXRates(coinCode, currencyCode, exchange)

		if err != nil {
			//TODO
		}
	}

	return (*LatestXRate)(data)
}

// Get method gets rates by Unix EPOCH date
func (xratesKit *XRatesKit) Get(
	coinCode string, currencyCode string, exchange string, epochSec int64) *HistoricalXRate {

	currencyCode = strings.ToUpper(currencyCode)
	coinCode = strings.ToUpper(coinCode)

	cacheData := xratesKit.cacheService.Get(coinCode, currencyCode, exchange, epochSec)

	if cacheData == nil {
		data, err := xratesKit.ipfsHandler.GetHistoricalXRates(coinCode, currencyCode, exchange, &epochSec)

		if err != nil {
			// Get data from CoinPaprika
			data, err = xratesKit.coinPaprika.GetHistoricalXRates(coinCode, currencyCode, exchange, &epochSec)

			if err != nil {

			}
		}

		//---------- Cache Data -------------------
		go xratesKit.cacheService.Set(coinCode, currencyCode, exchange, epochSec, data)
		log.Println("After set")
		//-----------------------------------------

		return (*HistoricalXRate)(data)

	} else {

		log.Println("Data from Cache")
		return (*HistoricalXRate)(cacheData)
	}
}
