package handler

import (
	"github.com/coinpaprika/coinpaprika-api-go-client/coinpaprika"
	"github.com/horizontalsystems/xrates-kit/pkg/config"
)

var coinpClient *coinpaprika.Client

//CoinPaprika handler object
type CoinPaprika struct {
	Conf *config.CoinPaprikaConfig
}

func init() {
	coinpClient = coinpaprika.NewClient(nil)
}

func (coinp CoinPaprika) GetLatestXRates(dCcy string, Ccy string, exchange string) (string, error) {
	return "", nil
}

func (coinp CoinPaprika) GetXRates(
	dCcy string, fCcy string, exchange string, epochSec *int64) (string, error) {

	return "", nil
}
