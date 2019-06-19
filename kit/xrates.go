package kit

import (
	"github.com/horizontalsystems/go-xrates-kit/cache"
	"github.com/horizontalsystems/go-xrates-kit/config"
	"github.com/horizontalsystems/go-xrates-kit/handler"
	"log"
	"strings"
	"time"
)

type XRatesKit struct {
	conf                     *config.MainConfig
	ipfsHandler, coinPaprika handler.XRates
	cacheService             *cache.CacheService
}

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

	return (*XRate)(xratesKit.cacheService.GetHistorical(coinCode, currencyCode, timestamp))
}

//GetLatest gets latest rates of source and target currencies
func (xratesKit *XRatesKit) GetLatest(currencyCode string, coinCodes *CoinCodes) *XRates {

	data, err := xratesKit.ipfsHandler.GetLatestXRates(currencyCode, coinCodes.toStringArray())

	if err != nil {

		// Get data from CoinPaprika
		data, err = xratesKit.coinPaprika.GetLatestXRates(currencyCode, coinCodes.toStringArray())

		if err != nil {
			return nil
		}
	}
	//---------- Cache Data -------------------

	xratesKit.cacheService.SetLatest(data)

	//-----------------------------------------

	log.Println(data)

	var result = make([]XRate, len(data))
	for i, xrate := range data {
		result[i] = XRate(xrate)
	}
	return &XRates{result}
}

//GetLatest gets latest rates of source and target currencies
func (xratesKit *XRatesKit) GetLatestCached(currencyCode string, coinCodes *CoinCodes) *XRates {

	currencyCode = strings.ToUpper(currencyCode)
	result := make([]XRate, coinCodes.Size())

	for i, coinCode := range coinCodes.toStringArray() {
		coinCode = strings.ToUpper(coinCode)
		xrate := xratesKit.cacheService.GetLatest(coinCode, currencyCode)
		result[i] = XRate(*xrate)
	}

	return &XRates{result}
}

//
////---------------------------Subscription------------------------------------

var latestRatesSyncer LatestRatesSyncer

type LatestRateListener interface {
	OnUpdate(rates *XRates)
}

func (xratesKit *XRatesKit) SubscribeToLatestRates(syncIntervalSecs int, currencyCode string, coinCodes *CoinCodes,
	listener LatestRateListener) {

	if latestRatesSyncer.started {
		latestRatesSyncer.Channel <- LatestRatesChannelItem{currencyCode, coinCodes, listener, syncIntervalSecs}
	} else {
		go latestRatesSyncer.start(xratesKit, syncIntervalSecs, currencyCode, coinCodes, listener)
	}

	//go startSyncerJob(xratesKit, syncIntervalSecs, currencyCode, coinCodes, listener)

}

type LatestRatesChannelItem struct {
	currencyCode          string
	coinCodes             *CoinCodes
	listener              LatestRateListener
	syncIntervalInSeconds int
}

type LatestRatesSyncer struct {
	started               bool
	currencyCode          string
	coinCodes             *CoinCodes
	listener              LatestRateListener
	syncIntervalInSeconds int
	Channel               chan LatestRatesChannelItem
}

func (s *LatestRatesSyncer) start(xratesKit *XRatesKit, syncIntervalSecs int, currencyCode string, coinCodes *CoinCodes,
	listener LatestRateListener) {

	latestRatesSyncer.started = true
	latestRatesSyncer.Channel = make(chan LatestRatesChannelItem)
	latestRatesSyncer.coinCodes = coinCodes
	latestRatesSyncer.currencyCode = currencyCode
	latestRatesSyncer.listener = listener
	latestRatesSyncer.syncIntervalInSeconds = syncIntervalSecs

	for {

		select {
		case channelItem := <-s.Channel:
			{
				log.Print("FROM CHANNEL ", channelItem)

				s.currencyCode = channelItem.currencyCode
				s.syncIntervalInSeconds = channelItem.syncIntervalInSeconds
				s.coinCodes = channelItem.coinCodes
				s.listener = channelItem.listener
			}

		default:
			{
				log.Print("DEFAULT")

				latestRates := xratesKit.GetLatest(s.currencyCode, s.coinCodes)

				log.Print("fetched latest rates", latestRates)

				if latestRates != nil {
					s.listener.OnUpdate(latestRates)
				}

				time.Sleep(time.Second * time.Duration(s.syncIntervalInSeconds))

				log.Print("waked up")
			}

		}
	}
}
