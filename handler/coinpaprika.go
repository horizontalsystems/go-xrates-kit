package handler

import (
	"fmt"
	"github.com/horizontalsystems/xrates-kit/models"
	"time"

	"github.com/coinpaprika/coinpaprika-api-go-client/coinpaprika"
	"github.com/horizontalsystems/xrates-kit/config"
)

var coinpClient *coinpaprika.Client

//CoinPaprika handler object
type CoinPaprika struct {
	Conf *config.CoinPaprikaConfig
}

func init() {
	coinpClient = coinpaprika.NewClient(nil)
}

func (cpHandler *CoinPaprika) GetLatestXRates(coinCode string, currencyCode string, exchange string) (
	*models.LatestXRate, error) {

	var err error
	var response *models.LatestXRate
	var index int

	t := time.Now().UTC()

	coin := models.COINS[coinCode]
	options := coinpaprika.TickersOptions{currencyCode}

	if coinCode == "" {

		for _, coin := range models.COINS {

			options := coinpaprika.TickersOptions{currencyCode}
			ticker, err := coinpClient.Tickers.GetByID(coin.GetFullName(), &options)
			if err != nil {

			}

			response = convertCoinpLatestData(response, ticker, coinCode, currencyCode, &t)

			// Coin paprika limits 10 request per second
			if index%8 == 0 && index != 0 {
				time.Sleep(1 * time.Second)
			}
			index++
		}
	} else {

		ticker, err := coinpClient.Tickers.GetByID(coin.GetFullName(), &options)
		if err != nil {
		}

		response = convertCoinpLatestData(response, ticker, coinCode, currencyCode, &t)

	}

	return response, err
}

func (coinp *CoinPaprika) GetHistoricalXRates(
	coinCode string, currencyCode string, exchange string, epochSec *int64) (*models.HistoricalXRate, error) {

	options := coinpaprika.TickersHistoricalOptions{}
	options.Start = time.Unix(*epochSec, 0).UTC()
	options.Quote = models.GetBaseCurreny().ID
	options.Limit = 1

	coin := models.COINS[coinCode]

	ticker, err := coinpClient.Tickers.GetHistoricalTickersByID(coin.GetFullName(), &options)

	if err != nil {

	}

	return convertCoinpHistoricalData(ticker, coinCode, currencyCode, &options.Start), err

}

func convertCoinpLatestData(xratesData *models.LatestXRate,
	ticker *coinpaprika.Ticker, coinCode string, currencyCode string, timeObj *time.Time) *models.LatestXRate {

	if xratesData == nil {
		xratesData = &models.LatestXRate{currencyCode, timeObj.String(), timeObj.Unix(), make(map[string]string)}
	}

	for _, element := range ticker.Quotes {

		xratesData.Rates[*ticker.Symbol] = fmt.Sprintf("%f", *element.Price)

		break
	}

	return xratesData
}

func convertCoinpHistoricalData(
	ticker []*coinpaprika.TickerHistorical,
	coinCode string, currencyCode string, timeObj *time.Time) *models.HistoricalXRate {

	for _, element := range ticker {

		baseCcy := models.GetBaseCurreny().ID
		tEpoch := timeObj.Unix()
		var fxRate float64 = 1

		if currencyCode != baseCcy {
			fxRes, _ := FiatXRatesHandler.GetXRates(baseCcy, currencyCode, "", &tEpoch)
			fxRate = fxRes.Rates[currencyCode]
		}

		fxRate = fxRate * (*element.Price)

		return &models.HistoricalXRate{coinCode, currencyCode, fmt.Sprintf("%f", fxRate), timeObj.Unix()}
	}

	return nil

}
