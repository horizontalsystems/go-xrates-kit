package handler

import (
	"fmt"
	"github.com/horizontalsystems/go-xrates-kit/models"
	"time"

	"github.com/coinpaprika/coinpaprika-api-go-client/coinpaprika"
	"github.com/horizontalsystems/go-xrates-kit/config"
)

var coinpClient *coinpaprika.Client

//CoinPaprika handler object
type CoinPaprika struct {
	Conf *config.CoinPaprikaConfig
}

func init() {
	coinpClient = coinpaprika.NewClient(nil)
}

func (cpHandler *CoinPaprika) GetLatestXRates(currencyCode string, coinCodes []string) ([]models.XRate, error) {

	var err error
	var index int

	t := time.Now().UTC()

	result := make([]models.XRate, len(coinCodes))

	for i, coinCode := range coinCodes {

		coin, ok := models.COINS[coinCode]
		if ok {

			options := coinpaprika.TickersOptions{currencyCode}
			ticker, err := coinpClient.Tickers.GetByID(coin.GetFullName(), &options)
			if err != nil {
				return nil, err
			} else {

				data, ok := convertCoinpLatestData(ticker, coinCode, currencyCode, &t)

				if ok {
					result[index] = *data
					index++
				}
			}
			// Coin paprika limits 10 request per second
			if i%8 == 0 && i != 0 {
				time.Sleep(1 * time.Second)
			}
		}
	}

	return result, err
}

func (coinp *CoinPaprika) GetHistoricalXRates(
	currencyCode string, coinCode string, timestamp *int64) (*models.XRate, error) {

	options := coinpaprika.TickersHistoricalOptions{}
	options.Start = time.Unix(*timestamp, 0).UTC()
	options.Quote = models.GetBaseCurreny().ID
	options.Limit = 1

	coin := models.COINS[coinCode]

	ticker, err := coinpClient.Tickers.GetHistoricalTickersByID(coin.GetFullName(), &options)

	if err != nil {

	}

	return convertCoinpHistoricalData(ticker, coinCode, currencyCode, &options.Start), err

}

func convertCoinpLatestData(ticker *coinpaprika.Ticker, coinCode string, currencyCode string, timeObj *time.Time) (*models.XRate, bool) {

	for _, element := range ticker.Quotes {

		return &models.XRate{coinCode, currencyCode, timeObj.Unix(), fmt.Sprintf("%f", *element.Price)}, true

	}

	return nil, false
}

func convertCoinpHistoricalData(ticker []*coinpaprika.TickerHistorical,
	coinCode string, currencyCode string, timeObj *time.Time) *models.XRate {

	for _, element := range ticker {

		baseCcy := models.GetBaseCurreny().ID
		tEpoch := timeObj.Unix()
		var fxRate float64 = 1

		if currencyCode != baseCcy {
			fxRes, _ := FiatXRatesHandler.GetXRates(baseCcy, currencyCode, "", &tEpoch)
			fxRate = fxRes.Rates[currencyCode]
		}

		fxRate = fxRate * (*element.Price)

		return &models.XRate{coinCode, currencyCode, timeObj.Unix(), fmt.Sprintf("%f", fxRate)}
	}

	return nil
}
