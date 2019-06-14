package handler

import (
	"encoding/json"
	"fmt"
	"github.com/horizontalsystems/go-xrates-kit/models"
	"strconv"
	"strings"
	"time"

	"github.com/horizontalsystems/go-xrates-kit/config"
	httputil "github.com/horizontalsystems/go-xrates-kit/util/http"
)

//Ipfs handler object
type Ipfs struct {
	Conf *config.IpfsConfig
}

func (ipfs *Ipfs) GetLatestXRates(coinCode string, currencyCode string, exchange string) (
	*models.LatestXRate, error) {

	respStr, err := httputil.DoGet(TIMEOUT_IPFS, ipfs.Conf.URL,
		"ipns/"+ipfs.Conf.IpnsID+"/xrates/latest/"+currencyCode+"/index.json", "")

	if err != nil {
		if strings.Contains(err.Error(), "Client.Timeout") {
			respStr, err = httputil.DoGet(TIMEOUT_IPFS, ipfs.Conf.PublicURL,
				"ipns/"+ipfs.Conf.IpnsID+"/xrates/latest/"+currencyCode+"/index.json", "")
		}
	}

	return reformatIPFSLatestData(respStr), err
}

func (ipfs *Ipfs) GetHistoricalXRates(coinCode string, currencyCode string, exchange string, epochSec *int64) (
	*models.HistoricalXRate, error) {

	var uriPathMinute, uriPathDay string

	timeSecObj := time.Unix(*epochSec, 0).UTC()
	year, month, day := timeSecObj.Date()
	hour := timeSecObj.Hour()

	uriPathDay = "ipns/" + ipfs.Conf.IpnsID +
		"/xrates/historical/" + coinCode + "/" + currencyCode + "/" +
		strconv.Itoa(year) + "/" +
		fmt.Sprintf("%02d", int(month)) + "/" +
		fmt.Sprintf("%02d", day)

	uriPathMinute = uriPathDay + "/" + fmt.Sprintf("%02d", hour) + "/index.json"

	respStr, err := httputil.DoGet(TIMEOUT_IPFS, ipfs.Conf.URL, uriPathMinute, "")

	if err != nil {

		if strings.Contains(err.Error(), "Client.Timeout") {
			respStr, err = httputil.DoGet(TIMEOUT_IPFS, ipfs.Conf.PublicURL, uriPathDay, "")

			if err != nil {
			}

		} else {
			// if Minute data is not found try By DAY.
			uriPathDay += "/index.json"

			respStr, err = httputil.DoGet(TIMEOUT_IPFS, ipfs.Conf.URL, uriPathDay, "")

			if err != nil {
			}
		}
	}

	return reformatIPFSHistoricalData(respStr, coinCode, currencyCode, &timeSecObj), err

}

//Change response json to XRates compatible format
func reformatIPFSLatestData(jsonData string) *models.LatestXRate {

	data := models.LatestXRate{}
	err := json.Unmarshal([]byte(jsonData), &data)

	if err != nil {

	}

	return &data
}

//Change response json to XRates compatible format
func reformatIPFSHistoricalData(jsonData string, coinCode string,
	currencyCode string, timeObj *time.Time) *models.HistoricalXRate {

	var rate string

	if jsonData == "" {
		return nil
	}

	data := map[string]string{}
	minute := fmt.Sprintf("%02d", timeObj.Minute())
	err := json.Unmarshal([]byte(jsonData), &data)

	if err != nil {
		rate = strings.ReplaceAll(jsonData, "\"", "")
	} else {
		rate = data[minute]
	}

	return &models.HistoricalXRate{coinCode, currencyCode, rate, timeObj.Unix()}
}
