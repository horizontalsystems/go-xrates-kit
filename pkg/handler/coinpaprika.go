package handler

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/coinpaprika/coinpaprika-api-go-client/coinpaprika"
	"github.com/horizontalsystems/xrates-kit/pkg/config"
	"github.com/horizontalsystems/xrates-kit/pkg/currency"
)

var coinpClient *coinpaprika.Client

//CoinPaprika handler object
type CoinPaprika struct {
	Conf *config.CoinPaprikaConfig
}

func init() {
	coinpClient = coinpaprika.NewClient(nil)
}

func (coinp *CoinPaprika) GetLatestXRatesAsJSON(dCcy string, fCcy string, exchange string) (string, error) {

	var err error
	var response *LatestXRatesData
	t := time.Now().UTC()

	options := coinpaprika.TickersOptions{fCcy}

	if dCcy == "" {

		for index, ccy := range currency.GetDigitalCurrencies() {

			options := coinpaprika.TickersOptions{fCcy}
			ticker, err := coinpClient.Tickers.GetByID(currency.GetFullName(ccy), &options)
			if err != nil {

			}

			response = convertCoinpLatestData(response, ticker, dCcy, fCcy, &t)

			// Coin paprika limits 10 request per second
			if index%8 == 0 && index != 0 {
				time.Sleep(1 * time.Second)
			}
		}
	} else {

		ticker, err := coinpClient.Tickers.GetByID(currency.GetFullName(dCcy), &options)
		if err != nil {
		}

		response = convertCoinpLatestData(response, ticker, dCcy, fCcy, &t)

	}

	bytes, err := json.Marshal(response)
	if err != nil {
	}

	return string(bytes), err
}

func (coinp *CoinPaprika) GetXRatesAsJSON(
	dCcy string, fCcy string, exchange string, epochSec *int64) (string, error) {

	options := coinpaprika.TickersHistoricalOptions{}
	options.Start = time.Unix(*epochSec, 0).UTC()
	options.Quote = currency.GetBaseFiatCurrency().ID
	options.Limit = 1

	ticker, err := coinpClient.Tickers.GetHistoricalTickersByID(currency.GetFullName(dCcy), &options)

	if err != nil {

	}

	return convertCoinpData(ticker, dCcy, fCcy, &options.Start), err

}

func convertCoinpLatestData(xratesData *LatestXRatesData,
	ticker *coinpaprika.Ticker, dCcy string, fCcy string, timeObj *time.Time) *LatestXRatesData {

	if xratesData == nil {
		xratesData = &LatestXRatesData{fCcy, timeObj.String(), timeObj.Unix(), make(map[string]string)}
	}

	for _, element := range ticker.Quotes {

		xratesData.Rates[*ticker.Symbol] = fmt.Sprintf("%f", *element.Price)

		break
	}

	return xratesData
}

func convertCoinpData(
	ticker []*coinpaprika.TickerHistorical, dCcy string, fCcy string, timeObj *time.Time) string {

	for _, element := range ticker {

		baseCcy := currency.GetBaseFiatCurrency().ID
		tEpoch := timeObj.Unix()
		var fxRate float64 = 1

		if fCcy != baseCcy {
			fxRes, _ := FiatXRatesHandler.GetXRates(baseCcy, fCcy, "", &tEpoch)
			fxRate = fxRes.Rates[fCcy]
		}

		fxRate = fxRate * (*element.Price)
		resp := map[string]XRatesData{}
		resp[dCcy] = XRatesData{fCcy, fmt.Sprintf("%f", fxRate), element.Timestamp.Unix()}

		bytes, err := json.Marshal(resp)
		if err != nil {
		}

		return string(bytes)
	}

	return ""
}
